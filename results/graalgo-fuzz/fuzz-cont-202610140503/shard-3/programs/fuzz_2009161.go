package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 int16; F2 string }
type GenS7 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{((("a" + "\n\t") + "hello") + ("\n\t" + ("a" + ("\n\t" + ("" + "foo bar"))))), "", ((("foo bar" + (("a" + "\n\t") + "a")) + ("foo bar" + "hello")) + "hello")}
    var v3 string = "foo bar"
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]string)
    v6 := []string{}
    for fi7 := 0; fi7 < 8; fi7++ {
        v3 = v3
        v3 = (v3 + v3)
        _ = fi7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int32 = 927
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: ("foo bar" + v3), F1: -950}
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 bool = false
            if v13 {
                v14 := GenS2{F0: 42}
                v3 = ""
                fmt.Println(fi12, fi11, v3)
                if (v13 || (v13 && (v13 && v13))) {
                    var v15 uint32 = 2
                    var v16 int8 = 1
                    _ = v15
                    _ = v16
                } else {
                    v17 := GenS3{F0: 2, F1: 0, F2: 4294967295}
                    _ = v17
                }
                _ = v14
            } else {
                v18 := GenS1{F0: ((v3 + "\n\t") + v3), F1: 0}
                v19 := GenS1{F0: v3, F1: 42}
                switch v3 {
                case "foo bar":
                    v3 = "a"
                case "a":
                    if ((!(0 > 0)) && ((1.5 != 1.5) && (100 > -71))) {
                        switch fi11 {
                        case 4294967295:
                            v3 = "a"
                            var v20 int64 = 880
                            var v21 uint32 = 668
                            _ = v20
                            _ = v21
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                case "hello":
                    switch fi12 {
                    case -293:
                        v22 := GenS4{F0: -1, F1: (("hello" + "hello") + v3)}
                        _ = v22
                    }
                }
                v23 := GenS4{F0: 34, F1: (((v3 + (v3 + "foo bar")) + v3) + (v3 + ("\n\t" + ("hello" + v3))))}
                _ = v18
                _ = v19
                _ = v23
            }
            fmt.Println(v13, v13, v3)
            _ = fi12
            _ = v13
        }
        fmt.Println(v3, fi11)
        _ = fi11
    }
    v24 := GenS1{F0: "", F1: 0}
    v3 = (v3 + "")
    if (996 == 1) {
        v3 = (((((v3 + ((("hello" + v3) + v3) + "")) + (v3 + ("hello" + v3))) + (v3 + v3)) + "foo bar") + ("hello" + (v3 + "a")))
    } else {
        if false {
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := GenS5{F0: 753}
                _ = fi25
                _ = v26
            }
            fmt.Println(v3)
        }
    }
    v3 = "\n\t"
    v27 := GenS4{F0: 1000, F1: v3}
    v28 := GenS6{F0: 2, F1: 1000, F2: ("foo bar" + ("a" + v3))}
    v3 = v3
    v3 = ""
    v29 := GenS3{F0: 65535, F1: 128, F2: 1}
    v30 := GenS7{F0: 127}
    v3 = ((("\n\t" + v3) + ("foo bar" + v3)) + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
