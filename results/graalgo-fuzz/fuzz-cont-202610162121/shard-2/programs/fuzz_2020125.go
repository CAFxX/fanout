package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int8 }
func genFn0() int {
    v6 := GenS1{F0: "a"}
    var v7 int64 = -9223372036854775808
    var v8 float64 = (1.5 + 0.1)
    v8 -= v8
    _ = v6
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 uint32, p1 int8, p2 int16) int16 {
    var v9 int64 = -1
    fmt.Println(v9, v9, p1)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p2
}

func main() {
    v1 := []uint32{}
    var v2 float64 = float64((0.0 - (1e10 + 3.14159)))
    var v3 uint = 325
    v4 := make(map[uint32]bool)
    var v5 string = ""
    fmt.Println(v3, v3)
    v2 = 0.1
    var v10 int = -1
    if (true || (true && false)) {
        for fi11 := 0; fi11 < 9; fi11++ {
            for fi12 := 0; fi12 < 14; fi12++ {
                fmt.Println(v2, v2, v2)
                var v13 uint = v3
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        v14 := GenS2{F0: -1, F1: 78}
        _ = v14
    } else {
        var v15 string = v5
        var v16 uint64 = 700
        _ = v15
        _ = v16
    }
    fmt.Println(v2)
    v10 += (-(v10))
    for fi17 := 0; fi17 < 9; fi17++ {
        switch ((("a" + (v5 + "\n\t")) + v5) + v5) {
        case "a":
            v2 = v2
        case "hello":
            var v18 uint64 = 932
            _ = v18
        default:
            switch (fi17 % ((fi17) | 1)) {
            case 4294967295:
                fmt.Println(v3, v10)
            }
        }
        _ = fi17
    }
    var v19 int = 32768
    var v20 int16 = -1
    v3 = (v3 * v3)
    var v21 int8 = -106
    switch (((v5 + "") + (("hello" + v5) + v5)) + v5) {
    case "\n\t":
        v22 := v4[2147483647]
        _ = v22
    default:
        v21 -= (v21 & v21)
    }
    var v23 uint16 = 221
    delete(v4, 65535)
    var v24 int8 = (v21 + (v21 * v21))
    if false {
        var v25 int16 = v20
        var v26 int = v10
        _ = v25
        _ = v26
    } else {
        if false {
            v24 = (v24 / ((-95) | 1))
            if true {
                var v27 int32 = -1
                if true {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        v29 := GenS3{F0: 255}
                        for fi30 := 0; fi30 < 2; fi30++ {
                            delete(v4, 0)
                            v23 += (v23 + v23)
                            v5 = v5
                            _ = fi30
                        }
                        var v31 uint32 = uint32(v3)
                        _ = fi28
                        _ = v29
                        _ = v31
                    }
                    for fi32 := 0; fi32 < 1; fi32++ {
                        v33 := GenS4{F0: 0, F1: 2, F2: (v24 + v24)}
                        v4[2147483648] = (!(((v23 * v23) < (v23 % (((v23 | v23)) | 1))) && (!false)))
                        if ((v3 >> ((18446744073709551615) & 63)) < 2) {
                            var v34 bool = ((float64(-2.25) + float64(1e10)) == v2)
                            _ = v34
                        } else {
                            fmt.Println(v19, v20, v20)
                            var v35 uint8 = 1
                            _ = v35
                        }
                        _ = fi32
                        _ = v33
                    }
                } else {
                    var v36 uint16 = 601
                    _ = v36
                }
                _ = v27
            } else {
                if ((65535 > 0) && false) {
                    var v37 uint16 = (v23 | v23)
                    fmt.Println(v3, v2, v20)
                    _ = v37
                } else {
                    var v38 int16 = v20
                    v1 = append(v1, 0)
                    _ = v38
                }
                if true {
                    fmt.Println(v24, v2, v10)
                }
                fmt.Println(v23)
            }
        } else {
            if true {
                var v39 float64 = v2
                var v40 int8 = (v24 >> ((v3) & 63))
                _ = v39
                _ = v40
            }
            fmt.Println(v10, v5, v24)
            fmt.Println(v19, v24, v24)
        }
        v5 = "hello"
    }
    v41 := GenS3{F0: 42}
    var v42 string = (v5 + (("hello" + v5) + v5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
