package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float64 = (-2.25 + (((1e10 - ((-2.25 / (-2.25 * 0.1)) + (0.1 + 0.0))) / 1.0) * float64((1.5 - 3.14159))))
    v2 := []int{-777, 742}
    var v3 uint8 = 127
    switch "" {
    case "":
        var v4 int32 = 840
        _ = v4
    case "a":
        if (false && ((float32(3.14159) != ((-2.25 + 1e10) * 3.14159)) || true)) {
            switch -548 {
            case 2147483647:
                if (!(!true)) {
                    var v5 int = 1
                    _ = v5
                }
            case 42:
                fmt.Println(v3)
            }
            if false {
                fmt.Println(v3, v3)
                v3 = v3
            }
            if (-9223372036854775808 >= -1) {
                for fi6 := 0; fi6 < 7; fi6++ {
                    var v7 uint64 = 1
                    for fi8 := 0; fi8 < 19; fi8++ {
                        var v9 uint16 = 32767
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v10 int = -274
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    _ = fi6
                    _ = v7
                }
                var v11 string = "hello"
                _ = v11
            }
        }
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            switch "" {
            case "a":
                if ((!false) && (32767 != 1)) {
                    var v13 float32 = (1e10 - 1e10)
                    _ = v13
                }
            }
            var v14 float64 = v1
            _ = fi12
            _ = v14
        }
    }
    var v15 int32 = 1
    if true {
        var v16 int16 = 32767
        v16 = 1
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 float32 = 0.0
        _ = fi17
        _ = v18
    }
    fmt.Println(v3)
    if (false || ((true && false) || (-1 <= 42))) {
        v1 = (v1 + float64((0.0 + 0.1)))
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v19 := GenS1{F0: 1000}
    var v20 int32 = (v15 & -845)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
