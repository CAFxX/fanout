package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint) bool {
    p0 -= 62
    v5 := []int{128, 32768, 1, 973}
    for fi6 := 0; fi6 < 4; fi6++ {
        if false {
            var v7 int32 = 42
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return (128 == 2)
}

func main() {
    var v1 uint = 0
    var v2 float64 = (3.14159 * 0.0)
    v3 := make(map[int]int)
    var v4 string = "a"
    v2 += float64(-2.25)
    var v8 int32 = -1
    v1 -= v1
    v9 := GenS1{F0: ("" + v4)}
    var v10 uint64 = 1
    var v11 uint8 = 128
    v2 = -2.25
    if false {
        if true {
            var v12 uint16 = 542
            v1 = (v1 - v1)
            _ = v12
        }
        if false {
            fmt.Println(v1)
        } else {
            var v13 string = v4
            _ = v13
        }
        var v14 int = 128
        _ = v14
    } else {
        _ = len(v3)
        v1 -= v1
        if (true || true) {
            v2 += ((v2 - 0.0) - v2)
            if (true || true) {
                v4 = v4
                for fi15 := 0; fi15 < 19; fi15++ {
                    switch "foo bar" {
                    case "hello":
                        v16 := []int{32768, (fi15 ^ (-(fi15)))}
                        _ = v16
                    case "foo bar":
                        v17 := v3[(fi15 * (fi15 / ((fi15) | 1)))]
                        _ = v17
                    case "a":
                        var v18 uint32 = 0
                        _ = v18
                    }
                    _ = fi15
                }
            } else {
                fmt.Println(v4, v4)
                for fi19 := 0; fi19 < 2; fi19++ {
                    var v20 string = v4
                    v11 = v11
                    for fi21 := 0; fi21 < 3; fi21++ {
                        v11 = (v11 >> ((1) & 63))
                        _ = fi21
                    }
                    _ = fi19
                    _ = v20
                }
                var v22 uint64 = v10
                _ = v22
            }
        }
    }
    var v23 int32 = v8
    v8 += -1
    v24 := GenS1{F0: (v4 + v4)}
    var v25 bool = (float64(float32(float64(1.5))) < v2)
    v4 = "\n\t"
    for fi26 := 0; fi26 < 20; fi26++ {
        for fi27 := 0; fi27 < 14; fi27++ {
            if (v11 >= v11) {
                var v28 uint8 = (v11 << (((v1 + (v1 * (v1 / ((4294967295) | 1))))) & 63))
                switch v4 {
                case "hello":
                    v29 := []int{}
                    _ = v29
                case "a":
                    fmt.Println(v10, v8, v1)
                case "\n\t":
                    var v30 uint8 = 46
                    _ = v30
                }
                var v31 int = (fi27 & fi27)
                _ = v28
                _ = v31
            }
            if (!(("\n\t" != ((("foo bar" + (v4 + v4)) + "") + v4)) && v25)) {
                switch v4 {
                case "foo bar":
                    v32 := []int{(fi26 << ((v1) & 63)), 937, fi26}
                    _ = v32
                }
            }
            _ = fi27
        }
        _ = fi26
    }
    delete(v3, 288)
    v33 := []int{int(v1), 1, 32768, int(v11)}
    v8 = v23
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v8
    _ = v9
}
