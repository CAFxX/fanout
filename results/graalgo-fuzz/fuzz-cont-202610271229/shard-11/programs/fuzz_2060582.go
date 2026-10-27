package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0() int8 {
    if (1 == 2) {
        v5 := []int{}
        v6 := GenS1{F0: -128}
        _ = v5
        _ = v6
    } else {
        v7 := []int{2147483647}
        if true {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 uint = 859
                fmt.Println(v9, v9)
                v10 := GenS1{F0: 127}
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 int16 = 42
            _ = v11
        }
        _ = v7
    }
    var v12 float64 = 3.14159
    for fi13 := 0; fi13 < 12; fi13++ {
        if true {
            var v14 uint32 = 2147483647
            switch (("\n\t" + "") + ("foo bar" + "foo bar")) {
            case "hello":
                var v15 uint8 = 1
                _ = v15
            case "a":
                v16 := []int{fi13, (fi13 / ((fi13) | 1))}
                _ = v16
            case "\n\t":
                switch fi13 {
                case 1:
                    v17 := GenS2{F0: 4294967295, F1: 1}
                    _ = v17
                case 128:
                    v18 := []int{fi13, (-(fi13)), fi13, (fi13 ^ fi13)}
                    _ = v18
                case 280:
                    v14 = v14
                }
            }
            _ = v14
        } else {
            var v19 string = ("foo bar" + "")
            _ = v19
        }
        _ = fi13
    }
    _ = v12
    return -128
}

func main() {
    var v1 float64 = (3.14159 - (0.0 * -2.25))
    var v2 float32 = float32(((1.5 / float64(1.5)) + float64(float32(float64((1e10 * (3.14159 - 0.1)))))))
    var v3 uint32 = 42
    var v4 uint64 = 39
    var v20 uint16 = 42
    v3 += (v3 ^ (v3 | (v3 / (((v3 ^ (v3 >> ((768) & 63)))) | 1))))
    fmt.Println(v3, v3)
    var v21 int32 = -613
    v21 += 42
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 int32 = v21
            fmt.Println(v4, v21)
            if true {
                v25 := []int{fi23, (fi23 & 0), (fi22 - (fi23 | 0))}
                v26 := GenS1{F0: 127}
                _ = v25
                _ = v26
            }
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v2 += float32(v1)
    v27 := []int{}
    v3 += 2147483647
    fmt.Println(v2, v3, v3)
    switch ("foo bar" + "hello") {
    case "a":
        if false {
            var v28 int8 = 1
            var v29 int8 = v28
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            _ = v28
            _ = v29
        } else {
            var v30 uint64 = (v4 >> ((1) & 63))
            v21 = 1
            _ = v30
        }
    case "\n\t":
        if len(v27) > 0 {
            v27[0] = 32767
        }
    case "hello":
        var v31 int64 = int64(v20)
        _ = v31
    }
    v4 = v4
    if false {
        var v32 int8 = -128
        for fi33 := 0; fi33 < 13; fi33++ {
            if true {
                if false {
                    v1 += v1
                }
            } else {
                for fi34 := 0; fi34 < 20; fi34++ {
                    var v35 int64 = 1
                    if (true || true) {
                        if true {
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            var v36 uint64 = (v4 - v4)
                            _ = v36
                        }
                        var v37 uint8 = 0
                        _ = v37
                    }
                    v38 := GenS1{F0: (-(v32))}
                    _ = fi34
                    _ = v35
                    _ = v38
                }
            }
            if ((1.5 + ((((v1 / (3.14159 - ((3.14159 / 1.0) + ((float64(v2) * float64(v2)) - v1)))) + 3.14159) * v1) * v1)) == (v1 - 1.5)) {
                if (((v32 << ((42) & 63)) > genFn0()) && false) {
                    var v39 uint64 = v4
                    _ = v39
                }
            } else {
                v21 = 377
                var v40 uint8 = 1
                _ = v40
            }
            _ = fi33
        }
        _ = v32
    }
    v4 -= 36
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
}
