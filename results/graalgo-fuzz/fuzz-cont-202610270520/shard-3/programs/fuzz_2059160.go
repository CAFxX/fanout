package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint; F1 uint8; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int }
func genFn0(p0 int) uint {
    v4 := []int{}
    if (!(false || false)) {
        if len(v4) > 0 {
            v4[0] = -563
        }
    } else {
        if true {
            var v5 uint32 = 2147483647
            v6 := GenS1{F0: 255}
            _ = v5
            _ = v6
        }
        var v7 int8 = 42
        _ = v7
    }
    var v8 uint64 = 1
    var v9 uint = 4294967295
    _ = p0
    _ = v4
    _ = v8
    _ = v9
    return v9
}
func genFn1(p0 uint, p1 uint, p2 uint16) uint32 {
    var v10 float64 = (1.5 * 1e10)
    var v11 bool = (p1 == (p0 >> ((p1) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 666
}

func main() {
    var v1 int32 = 0
    v2 := []int{32768, 0, 4294967295, 32767}
    var v3 uint32 = 2
    var v12 uint64 = 42
    switch 168 {
    case -184:
        _ = len(v2)
    case 1:
        if true {
            var v13 string = "hello"
            var v14 float64 = 1e10
            v15 := GenS1{F0: 0}
            _ = v13
            _ = v14
            _ = v15
        }
    }
    var v16 uint8 = 255
    var v17 int64 = 42
    var v18 uint = 255
    var v19 int16 = 32767
    v20 := GenS2{F0: 0, F1: (v12 * (v12 - (v12 + 2))), F2: 127}
    v21 := GenS3{F0: genFn0(-139), F1: v16, F2: 2}
    switch (("\n\t" + (("\n\t" + ("a" + "hello")) + "foo bar")) + ("" + (("hello" + "\n\t") + ("a" + "a")))) {
    case "":
        if (!false) {
            var v22 uint64 = (v12 / ((2) | 1))
            if ((-885 <= 4294967295) && true) {
                if false {
                    var v23 int64 = v17
                    v24 := GenS4{F0: 42}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v23
                    _ = v24
                }
                v19 += 940
                _ = len(v2)
            }
            v25 := GenS4{F0: v16}
            _ = v22
            _ = v25
        } else {
            var v26 float64 = 3.14159
            v16 = (v16 << ((536) & 63))
            v27 := GenS5{F0: (-(v19)), F1: -1}
            _ = v26
            _ = v27
        }
    case "a":
        if true {
            if false {
                v19 = 42
                var v28 int = -1
                _ = v28
            } else {
                if len(v2) > 0 {
                    v2[0] = int(v19)
                }
                if len(v2) > 0 {
                    v2[0] = -336
                }
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                v19 = 402
                if ((v12 & (v12 ^ (v12 * v12))) <= (v12 >> ((uint(v16)) & 63))) {
                    switch fi29 {
                    case 4294967295:
                        v16 = v16
                    case 205:
                        for fi30 := 0; fi30 < 14; fi30++ {
                            var v31 float32 = (float32(float64(-2.25)) - 0.1)
                            v19 -= v19
                            v16 += v16
                            _ = fi30
                            _ = v31
                        }
                    case 32767:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    switch (fi29 - -9223372036854775808) {
                    case 890:
                        if true {
                            if len(v2) > 0 {
                                v2[0] = fi29
                            }
                            var v32 uint64 = (v12 * v12)
                            _ = v32
                        }
                    }
                    var v33 int64 = v17
                    _ = v33
                }
                var v34 int32 = (-(v1))
                _ = fi29
                _ = v34
            }
            var v35 float32 = float32((0.1 / 1e10))
            _ = v35
        } else {
            v36 := GenS1{F0: 937}
            var v37 bool = ((v18 / ((1) | 1)) <= (^(v18)))
            _ = v36
            _ = v37
        }
    }
    switch "\n\t" {
    case "\n\t":
        var v38 string = "a"
        _ = v38
    case "a":
        var v39 uint = 255
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
