package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 bool, p1 int64) bool {
    var v4 uint8 = 245
    var v5 int64 = (p1 | (p1 << ((42) & 63)))
    v6 := GenS1{F0: 1, F1: "hello", F2: 1}
    v7 := GenS2{F0: 517, F1: uint32(v4)}
    var v8 int64 = v5
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1() uint {
    v9 := []int{90, -2147483648, 4294967295}
    var v10 bool = false
    if v10 {
        v10 = v10
    } else {
        v10 = (32767 > 0)
        switch ("\n\t" + ("a" + "")) {
        case "foo bar":
            for fi11 := 0; fi11 < 4; fi11++ {
                if (float64((1.5 * (1e10 / 1.5))) >= (-2.25 - -2.25)) {
                    _ = len(v9)
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v10 = (2147483647 == -2147483648)
                        v10 = v10
                        _ = fi12
                    }
                }
                for fi13 := 0; fi13 < 7; fi13++ {
                    fmt.Println(v10, v10)
                    v10 = v10
                    v10 = v10
                    switch "foo bar" {
                    case "hello":
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    case "a":
                        v9 = append(v9, (fi11 % (((fi13 & fi11)) | 1)))
                    case "":
                        var v14 int32 = 0
                        _ = v14
                    default:
                        var v15 uint16 = 317
                        _ = v15
                    }
                    _ = len(v9)
                    _ = fi13
                }
                _ = fi11
            }
        case "hello":
            var v16 string = "\n\t"
            _ = v16
        default:
            switch 903 {
            case 127:
                var v17 int = 461
                _ = v17
            }
        }
    }
    var v18 int64 = 0
    _ = v10
    _ = v18
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    if ((-456 < -32768) || true) {
        switch "\n\t" {
        case "foo bar":
            var v19 uint16 = 65535
            _ = v19
        case "hello":
            v20 := GenS2{F0: -230, F1: 861}
            _ = v20
        }
    } else {
        var v21 uint = 134
        var v22 uint = 255
        var v23 float32 = (1e10 - (0.0 - float32(float64(1e10))))
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 bool = ((!false) || (2 <= 0))
    fmt.Println(v24)
    v3[32768] = (100 < 60)
    fmt.Println(v24)
    var v25 uint = 0
    v26 := GenS1{F0: 293, F1: "", F2: 4294967295}
    var v27 int64 = -6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
