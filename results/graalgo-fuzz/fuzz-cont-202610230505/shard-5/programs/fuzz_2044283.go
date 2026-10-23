package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int8) uint {
    v7 := []int{}
    v8 := GenS1{F0: 2}
    var v9 int64 = 376
    fmt.Println(v9, p0)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int8, p1 string, p2 bool) uint8 {
    var v10 uint32 = 0
    p2 = (!(!false))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 224
    var v3 uint = 65535
    var v4 int32 = -949
    v5 := make(map[int]uint32)
    v6 := make(map[string]int)
    if true {
        v11 := GenS1{F0: (v2 + v2)}
        var v12 int64 = 1
        _ = v11
        _ = v12
    }
    var v13 int16 = 402
    if ((v13 / (((v13 * (v13 | 426))) | 1)) <= -1) {
        if (3.14159 < 3.14159) {
            switch (((("a" + "") + "a") + ("a" + "a")) + (((("" + "") + ((("foo bar" + ("" + "\n\t")) + "foo bar") + ("hello" + ""))) + "\n\t") + "a")) {
            case "":
                v13 = -32768
            }
        }
        if (true || false) {
            v2 += 14
            v14 := GenS2{F0: 1}
            fmt.Println(v13, v3, v2)
            _ = v14
        } else {
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 int16 = (^(v13))
                var v17 uint64 = (v2 << ((v3) & 63))
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 uint16 = 571
            _ = v18
        }
        var v19 int64 = 1
        _ = v19
    }
    v20 := GenS2{F0: 0}
    var v21 float32 = (1.5 - ((-2.25 * ((1.5 * 1.5) + 0.0)) + 1.5))
    for fi22 := 0; fi22 < 19; fi22++ {
        v21 += ((v21 * 1e10) + ((v21 / v21) / ((v21 - (0.1 / (v21 - (v21 - v21)))) - (v21 + ((v21 - (1e10 * float32(float64(0.1)))) + v21)))))
        _ = fi22
    }
    if false {
        if false {
            if (((1.5 > 1.5) && false) || false) {
                for fi23 := 0; fi23 < 9; fi23++ {
                    v13 -= v13
                    var v24 uint8 = 128
                    _ = fi23
                    _ = v24
                }
                v2 = (v2 / ((657) | 1))
                fmt.Println(v3, v2)
            }
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            fmt.Println(v3, v2, v21)
            _ = fi25
        }
    } else {
        v3 = genFn0(-56)
        if true {
            if (v2 >= 1) {
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 bool = false
                    var v28 int = 32768
                    if v27 {
                        var v29 uint16 = 612
                        v6["a"] = v28
                        for fi30 := 0; fi30 < 5; fi30++ {
                            var v31 uint8 = genFn1(108, "foo bar", true)
                            v3 = v3
                            v3 = v3
                            _ = fi30
                            _ = v31
                        }
                        _ = v29
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            v32 := GenS1{F0: (v2 >> ((595) & 63))}
            var v33 int64 = 589
            _ = v32
            _ = v33
        }
    }
    v4 = int32(v2)
    fmt.Println(v2)
    v2 -= 939
    for fi34 := 0; fi34 < 12; fi34++ {
        v21 = float32(0.1)
        _ = fi34
    }
    v35 := GenS2{F0: 1}
    v36 := GenS3{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
