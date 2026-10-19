package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }
type GenS3 struct { F0 int32 }
func genFn0() int16 {
    var v6 uint64 = 41
    for fi7 := 0; fi7 < 14; fi7++ {
        switch (fi7 - fi7) {
        case 128:
            if (((0.1 / 1e10) < 0.0) || true) {
                fmt.Println(fi7, v6)
                var v8 uint64 = (v6 & v6)
                _ = v8
            }
        case 898:
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := GenS1{F0: 42, F1: 11}
                v6 += v6
                switch fi7 {
                case 32768:
                    v11 := []int{}
                    _ = v11
                default:
                    if true {
                        v6 -= v6
                        var v12 uint = 255
                        var v13 int16 = -32768
                        var v14 float64 = ((3.14159 * float64((0.0 - 1e10))) + 1e10)
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
                _ = fi9
                _ = v10
            }
        case 127:
            v15 := []int{128}
            _ = v15
        }
        fmt.Println(fi7)
        _ = fi7
    }
    _ = v6
    return -579
}
func genFn1(p0 bool, p1 uint16) uint64 {
    var v16 int = 463
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint = 18446744073709551615
        var v19 float32 = 1e10
        p1 -= p1
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: 1000, F1: 589}
    var v21 string = ("a" + ((("\n\t" + "\n\t") + "foo bar") + "a"))
    var v22 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v16
    _ = v20
    _ = v21
    _ = v22
    return 1
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[int]uint32)
    var v3 int64 = 42
    v4 := []int{32768}
    var v5 float64 = float64(0.1)
    v2[94] = 0
    if (255 >= 85) {
        var v23 string = ""
        if len(v4) > 0 {
            v4[0] = -905
        }
        if false {
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 uint64 = genFn1(true, 313)
                var v26 bool = (527 >= 2147483647)
                if (v3 <= (v3 * -1)) {
                    var v27 uint16 = 42
                    v23 = (v23 + v23)
                    v28 := v2[fi24]
                    _ = v27
                    _ = v28
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            v2[4294967295] = 42
            if true {
                if false {
                    v29 := GenS1{F0: genFn0(), F1: 610}
                    _ = v29
                } else {
                    fmt.Println(v23)
                    for fi30 := 0; fi30 < 5; fi30++ {
                        switch v23 {
                        case "hello":
                            var v31 int16 = 818
                            _ = v31
                        }
                        var v32 uint32 = 4294967295
                        v33 := v2[fi30]
                        _ = fi30
                        _ = v32
                        _ = v33
                    }
                    v1 -= v5
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v34 := GenS2{F0: true, F1: 0, F2: 65535}
                _ = v34
            }
        }
        _ = v23
    } else {
        if (42 < 2147483647) {
            var v35 string = (("\n\t" + (("a" + ("hello" + (("\n\t" + "foo bar") + "hello"))) + "foo bar")) + "foo bar")
            _ = v35
        } else {
            v36 := GenS1{F0: 176, F1: 42}
            _ = v36
        }
        if (65535 == 353) {
            if true {
                _ = len(v4)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = len(v4)
        }
    }
    v37 := GenS3{F0: 1}
    v38 := GenS1{F0: 1, F1: 1}
    for fi39 := 0; fi39 < 20; fi39++ {
        fmt.Println(v5, v1, v1)
        v3 += 42
        v3 = (v3 / (((-(v3))) | 1))
        _ = fi39
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        fmt.Println(v1, v3)
        delete(v2, (fi40 >> ((206) & 63)))
        _ = fi40
    }
    if (true && (true && true)) {
        v3 = -358
        switch 128 {
        case 409:
            v2[4294967295] = 986
        case -1:
            fmt.Println(v5, v1)
        case 32767:
            for fi41 := 0; fi41 < 6; fi41++ {
                var v42 uint = 2
                v42 -= (v42 / ((255) | 1))
                _ = fi41
                _ = v42
            }
        }
        v3 -= v3
    }
    if len(v4) > 0 {
        v4[0] = 32767
    }
    if false {
        var v43 int32 = -1
        v44 := GenS3{F0: v43}
        var v45 int8 = 100
        _ = v43
        _ = v44
        _ = v45
    }
    v46 := GenS2{F0: true, F1: 147, F2: 42}
    delete(v2, 32767)
    v5 -= (v5 - ((((0.1 - (v5 - 0.0)) / ((v5 + v5) * 1.5)) - 3.14159) - 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v46
    _ = v5
}
