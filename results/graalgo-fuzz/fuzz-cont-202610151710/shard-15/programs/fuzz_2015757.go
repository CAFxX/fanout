package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool; F1 int8; F2 string }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 string; F1 uint8; F2 string }

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]bool)
    var v3 uint64 = 42
    fmt.Println(v1, v3)
    v4 := []int{115, -1, 32768}
    fmt.Println(v1)
    v1 += 0
    if false {
        if true {
            var v5 string = "hello"
            if (686 != 42) {
                switch "a" {
                case "hello":
                    for fi6 := 0; fi6 < 2; fi6++ {
                        if ((!((v5 + v5) >= v5)) || (!(v3 == v3))) {
                            v3 = 42
                            v7 := GenS1{F0: 1, F1: 2}
                            var v8 uint64 = v3
                            _ = v7
                            _ = v8
                        }
                        v9 := GenS2{F0: (3.14159 == ((((1.5 * -2.25) * -2.25) - 0.1) - 1e10)), F1: -1, F2: (v5 + v5)}
                        _ = fi6
                        _ = v9
                    }
                }
                for fi10 := 0; fi10 < 6; fi10++ {
                    fmt.Println(v5)
                    fmt.Println(v5, v5)
                    _ = len(v2)
                    _ = fi10
                }
            }
            var v11 uint32 = 4294967295
            _ = v11
            _ = v5
        } else {
            v12 := GenS2{F0: (!(32767 >= 32767)), F1: 118, F2: ("a" + "hello")}
            var v13 int64 = 0
            v2[("foo bar" + (((((("hello" + "foo bar") + (("a" + "\n\t") + ("foo bar" + "\n\t"))) + "a") + "\n\t") + "\n\t") + "foo bar"))] = true
            _ = v12
            _ = v13
        }
        if true {
            var v14 uint32 = 2
            for fi15 := 0; fi15 < 19; fi15++ {
                for fi16 := 0; fi16 < 6; fi16++ {
                    if false {
                        if true {
                            var v17 uint32 = (v14 ^ 624)
                            v3 = v3
                            _ = v17
                        }
                        v18 := GenS1{F0: 32767, F1: 31}
                        _ = v18
                    }
                    switch 42 {
                    case 2147483647:
                        v3 = (v3 | 0)
                    case 0:
                        v3 += (v3 >> ((42) & 63))
                    case 128:
                        var v19 string = "\n\t"
                        _ = v19
                    }
                    _ = fi16
                }
                _ = fi15
            }
            _ = v14
        } else {
            v20 := GenS3{F0: 0, F1: 42, F2: 42}
            _ = v20
        }
        var v21 uint16 = uint16(v3)
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        fmt.Println(fi22)
        v23 := GenS1{F0: 0, F1: 162}
        for fi24 := 0; fi24 < 5; fi24++ {
            var v25 uint16 = 32767
            fmt.Println(v1, v1, v25)
            var v26 bool = true
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi22
        _ = v23
    }
    v27 := GenS1{F0: 65535, F1: 1}
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 int = (fi28 % ((-807) | 1))
        switch "\n\t" {
        case "a":
            for fi30 := 0; fi30 < 6; fi30++ {
                v2[((("" + "") + "") + "foo bar")] = false
                v31 := GenS4{F0: ("hello" + ("" + "\n\t")), F1: 2, F2: "hello"}
                var v32 float32 = (3.14159 - 0.0)
                _ = fi30
                _ = v31
                _ = v32
            }
        case "foo bar":
            _ = len(v4)
        case "hello":
            fmt.Println(v29, v29, v3)
        }
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
