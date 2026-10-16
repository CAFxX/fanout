package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int32; F1 bool; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int16 }
func genFn0() uint32 {
    var v7 int64 = 42
    switch 0 {
    case -347:
        var v8 int8 = 127
        _ = v8
    case -2147483648:
        if (34 > 1) {
            var v9 uint16 = 674
            _ = v9
        } else {
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 uint64 = 42
                var v12 uint8 = 202
                _ = fi10
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 uint64 = 42
                v15 := GenS1{F0: 577, F1: 150, F2: 448}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
    }
    v16 := []int{}
    switch "\n\t" {
    case "":
        var v17 float64 = -2.25
        _ = v17
    default:
        var v18 int16 = -594
        _ = v18
    }
    var v19 int16 = -1
    _ = v16
    _ = v19
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"hello", "hello"}
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 int8 = -19
    var v6 int64 = 1
    v6 = v6
    var v20 int64 = v6
    if ("" > (("hello" + ("\n\t" + ("a" + "foo bar"))) + "foo bar")) {
        var v21 float64 = 1e10
        if false {
            fmt.Println(v20, v5, v5)
            if (!(0 >= 65535)) {
                var v22 uint8 = 255
                _ = v22
            }
            v23 := GenS2{F0: 0, F1: (837 <= 32767), F2: 42}
            _ = v23
        }
        _ = v21
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            if false {
                v5 += v5
                if ((v5 >> ((150) & 63)) >= v5) {
                    var v24 uint = 963
                    _ = v24
                }
                for fi25 := 0; fi25 < 6; fi25++ {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        v6 = -1
                        var v27 int = (fi26 >> ((638) & 63))
                        _ = fi26
                        _ = v27
                    }
                    _ = fi25
                }
            }
            _ = len(v1)
        }
        v28 := GenS1{F0: 2, F1: 32767, F2: 1}
        _ = v28
    }
    v6 = 9223372036854775807
    fmt.Println(v5, v20)
    v29 := GenS2{F0: 0, F1: (!true), F2: -137}
    v6 = (v20 - (v20 - v6))
    for fi30 := 0; fi30 < 19; fi30++ {
        for fi31 := 0; fi31 < 20; fi31++ {
            v5 = -47
            if true {
                fmt.Println(fi31, v6, v5)
            }
            var v32 uint = 65535
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    var v33 uint16 = 65535
    v34 := GenS2{F0: 2147483647, F1: (!true), F2: 0}
    v20 = v20
    v35 := GenS3{F0: (^(v33))}
    switch "\n\t" {
    case "foo bar":
        switch -769 {
        case 4294967295:
            if (!(3.14159 > 0.0)) {
                v36 := GenS4{F0: -1, F1: 0}
                for fi37 := 0; fi37 < 7; fi37++ {
                    var v38 int32 = 0
                    switch "hello" {
                    case "foo bar":
                        v39 := GenS3{F0: v33}
                        _ = v39
                    case "hello":
                        v40 := v4[""]
                        _ = v40
                    default:
                        if (-1 >= -646) {
                            var v41 int = (fi37 ^ fi37)
                            _ = v41
                        }
                    }
                    _ = len(v2)
                    _ = fi37
                    _ = v38
                }
                _ = v36
            }
        case 486:
            fmt.Println(v33, v33)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
