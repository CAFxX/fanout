package main

import (
    "fmt"
)

func genFn0(p0 int64, p1 int32, p2 uint8) uint {
    v7 := []int{}
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(p2, p2, p0)
        p1 -= (p1 ^ p1)
        p1 -= (p1 + (p1 | -1))
        _ = fi8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(p0, p0)
        for fi10 := 0; fi10 < 5; fi10++ {
            p2 -= 171
            var v11 uint64 = 1
            if (true || true) {
                fmt.Println(p0)
            }
            _ = fi10
            _ = v11
        }
        var v12 uint = 18446744073709551615
        _ = fi9
        _ = v12
    }
    var v13 int8 = 90
    var v14 float64 = (((float64((((-2.25 * float32(3.14159)) - 0.0) - float32(-2.25))) * (-2.25 * (0.1 - (0.1 * (3.14159 + (1e10 + -2.25)))))) * (1.5 - 1e10)) / (float64(3.14159) * 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v7
    return 0
}
func genFn1(p0 string) string {
    if (true && true) {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 float64 = 0.1
            p0 = p0
            _ = fi15
            _ = v16
        }
    }
    var v17 uint16 = 423
    _ = p0
    _ = v17
    return p0
}

func main() {
    var v1 uint8 = 2
    var v2 string = ("\n\t" + ("\n\t" + "foo bar"))
    var v3 uint64 = 0
    var v4 float64 = 0.0
    v5 := []int{-1, 32768, -9223372036854775808, 42}
    var v6 uint = 18446744073709551615
    if (469 < 65535) {
        if true {
            v1 = v1
            switch 885 {
            case 4294967295:
                if (!true) {
                    var v18 uint = genFn0(0, -651, (^(v1)))
                    if true {
                        var v19 uint64 = v3
                        _ = v19
                    }
                    _ = v18
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v4 -= float64(-2.25)
                }
            case -514:
                if len(v5) > 0 {
                    v5[0] = -831
                }
            case 32767:
                var v20 uint64 = (v3 - (^(v3)))
                _ = v20
            }
            v6 += genFn0(-80, -1, v1)
        }
    } else {
        var v21 bool = true
        v4 -= v4
        _ = v21
    }
    v1 += (v1 | (^(v1)))
    var v22 uint = (v6 * (^(v6)))
    var v23 int32 = -344
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v6)
    if false {
        for fi24 := 0; fi24 < 6; fi24++ {
            var v25 float32 = 1e10
            v23 = v23
            _ = fi24
            _ = v25
        }
    }
    v3 = 559
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
