package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint8) uint {
    p0 = 1
    var v4 uint64 = 232
    fmt.Println(p0, v4)
    _ = p0
    _ = v4
    return 14
}

func main() {
    var v1 uint8 = 127
    var v2 float32 = ((((0.0 * 0.1) * 3.14159) + 0.0) + ((float32(1e10) - 0.1) - float32(3.14159)))
    var v3 float64 = 0.1
    var v5 int64 = 0
    v1 += (v1 - 42)
    var v6 int = 2147483647
    var v7 int = 1
    v1 += 127
    if (v3 >= (0.1 * -2.25)) {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 uint8 = 0
            v7 = -773
            v6 = (^(fi8))
            _ = fi8
            _ = v9
        }
        v3 = (v3 / (float64((((float32(((v3 - (((v3 + -2.25) - ((v3 + -2.25) - v3)) - 0.0)) / v3)) / -2.25) / v2) + (((((1e10 + float32(v3)) - float32(v3)) / v2) / ((v2 * 1e10) + v2)) + 0.0))) - (float64(1e10) / v3)))
    }
    var v10 string = ("a" + "a")
    v11 := []int{v6, 907}
    v5 = (v5 / (((^(v5))) | 1))
    var v12 uint64 = 859
    switch v10 {
    case "":
        v13 := GenS1{F0: 871}
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 int = v7
            _ = fi14
            _ = v15
        }
    }
    var v16 string = ("" + (v10 + ((v10 + v10) + "")))
    var v17 int = v6
    for fi18 := 0; fi18 < 19; fi18++ {
        if true {
            v3 += 0.0
            fmt.Println(v17)
        }
        if false {
            var v19 int32 = 0
            if (v19 <= 992) {
                fmt.Println(fi18)
            } else {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                switch -635 {
                case 604:
                    var v20 int16 = 42
                    _ = v20
                case 32767:
                    v1 += v1
                }
                fmt.Println(v19, v1, v12)
            }
            var v21 uint64 = v12
            _ = v19
            _ = v21
        } else {
            fmt.Println(v7, fi18, v17)
        }
        if false {
            if true {
                v3 -= 0.1
                v1 += (v1 << ((2) & 63))
                v3 = ((float64(1e10) - 3.14159) * ((float64(v2) / float64((((v2 + v2) - float32(v3)) - float32((1.5 / ((1e10 * ((v3 + v3) + v3)) + ((float64(((((((v2 - v2) * (((v2 * v2) / (-2.25 / (v2 + ((v2 - v2) + (float32(v3) * v2))))) / 1.5)) * (v2 / (v2 - v2))) - (0.0 - float32(v3))) * ((3.14159 / float32((1e10 / v3))) + ((v2 - 3.14159) * float32(3.14159)))) + 1e10)) - (((float64(v2) + (v3 + (float64(1e10) / (v3 + (v3 - (v3 - 0.1)))))) + (v3 - v3)) / 3.14159)) + (((float64(float32(((1.5 * 0.0) + -2.25))) + ((v3 * (v3 + v3)) - (v3 + v3))) + ((v3 + 1.5) + v3)) + (1.5 + 0.1))))))))) + v3))
            }
            v7 += (-(v17))
            var v22 bool = (!((v3 / (1e10 - v3)) >= v3))
            _ = v22
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
