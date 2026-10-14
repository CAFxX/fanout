package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint64 }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0(p0 int64) uint {
    if (328 != 1) {
        p0 = p0
    } else {
        var v4 uint8 = 128
        _ = v4
    }
    if true {
        v5 := GenS1{F0: -123, F1: true, F2: 18446744073709551615}
        _ = v5
    } else {
        v6 := []int{}
        p0 -= p0
        _ = v6
    }
    var v7 int32 = -2147483648
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]string)
    var v3 uint16 = 255
    v1 = (^(v1))
    v3 -= v3
    v8 := GenS1{F0: -1, F1: (v1 == v1), F2: 2}
    for fi9 := 0; fi9 < 8; fi9++ {
        if (fi9 == (fi9 % (((fi9 + 405)) | 1))) {
            if (("\n\t" + "\n\t") < ("foo bar" + "hello")) {
                v10 := GenS2{F0: true, F1: 648}
                switch fi9 {
                case 0:
                    for fi11 := 0; fi11 < 4; fi11++ {
                        v12 := []int{1, 127, 32767}
                        if (-1 == -2147483648) {
                            v13 := GenS1{F0: 25, F1: false, F2: 87}
                            v3 = (v3 / ((v3) | 1))
                            _ = v13
                        }
                        _ = fi11
                        _ = v12
                    }
                case 1:
                    var v14 float32 = 1.5
                    _ = v14
                case 137:
                    switch (fi9 % (((fi9 * (fi9 ^ fi9))) | 1)) {
                    case 42:
                        _ = len(v2)
                    }
                }
                v15 := GenS2{F0: (!(!(255 >= 0))), F1: 9223372036854775807}
                _ = v10
                _ = v15
            }
        }
        var v16 int32 = 0
        _ = fi9
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint64 = 4294967295
        for fi19 := 0; fi19 < 5; fi19++ {
            if false {
                v18 = (^(v18))
            }
            if false {
                fmt.Println(v18, fi17)
                if true {
                    var v20 uint = 1
                    _ = v20
                } else {
                    fmt.Println(v18)
                    fmt.Println(fi19, v1, fi19)
                }
                if false {
                    if false {
                        v18 += 228
                        fmt.Println(fi19, fi19, v3)
                        v18 += v18
                    } else {
                        if false {
                            var v21 uint64 = (v18 >> ((34) & 63))
                            var v22 int8 = int8(v1)
                            _ = v21
                            _ = v22
                        } else {
                            var v23 int16 = -932
                            v24 := GenS1{F0: int8(fi17), F1: true, F2: v18}
                            _ = v23
                            _ = v24
                        }
                        v25 := []int{fi19, fi19, (fi17 >> ((255) & 63)), 42}
                        v26 := v2[65535]
                        fmt.Println(v18)
                        _ = v25
                        _ = v26
                    }
                    var v27 uint32 = 2147483648
                    _ = v27
                }
            } else {
                if (!false) {
                    var v28 int64 = -1
                    fmt.Println(fi17, fi19)
                    var v29 uint16 = (^(v3))
                    _ = v28
                    _ = v29
                }
            }
            v1 = v1
            v18 += v18
            var v30 uint32 = 2
            _ = fi19
            _ = v30
        }
        if ((((3.14159 - -2.25) - (1e10 + 0.1)) < -2.25) || true) {
            v31 := []int{fi17, (fi17 | fi17)}
            v32 := GenS1{F0: 100, F1: ((false && true) && (false || (!((("\n\t" + "foo bar") < "a") || false)))), F2: (v18 + 2)}
            var v33 uint8 = 42
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi17
        _ = v18
    }
    v34 := []int{1}
    v3 = v3
    v3 = (v3 * (v3 - 65535))
    var v35 string = ((((("foo bar" + "") + "hello") + "foo bar") + "hello") + "a")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v8
}
