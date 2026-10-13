package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint16, p1 int32, p2 int64) int8 {
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS1{F0: true, F1: "", F2: 18446744073709551615}
        p1 = (p1 + (p1 & 1))
        for fi7 := 0; fi7 < 6; fi7++ {
            fmt.Println(p2)
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    if (2 != 2) {
        fmt.Println(p0, p2, p2)
        var v8 int8 = int8(p2)
        _ = v8
    }
    var v9 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{}
    var v3 int64 = 42
    var v4 int8 = 100
    var v10 int32 = -1
    var v11 uint = 2
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v1, fi12, v1)
        switch (("a" + "hello") + "") {
        case "\n\t":
            var v13 int = 9223372036854775807
            _ = v13
        case "":
            if (945 != 0) {
                var v14 string = ((("\n\t" + ("hello" + (("foo bar" + "a") + "a"))) + "a") + ("a" + "a"))
                v15 := GenS2{F0: 0}
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 uint16 = 2
                    var v18 int32 = v10
                    var v19 string = "\n\t"
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                _ = v14
                _ = v15
            } else {
                if len(v2) > 0 {
                    v2[0] = 0
                }
                v4 = v4
                v20 := GenS2{F0: v10}
                _ = v20
            }
        }
        _ = fi12
    }
    v21 := GenS3{F0: v3}
    if (2147483647 > 4294967295) {
        v1 += 1e10
        v1 = 3.14159
    } else {
        _ = len(v2)
        v1 = ((0.1 + (0.0 - float32(0.1))) + (0.0 / 1.0))
    }
    v11 = 18446744073709551615
    v22 := GenS4{F0: 2, F1: (v11 | 42), F2: 0}
    v2 = append(v2, 56)
    if true {
        for fi23 := 0; fi23 < 19; fi23++ {
            v24 := GenS2{F0: v10}
            v4 = v4
            switch ((("foo bar" + (("a" + "") + "a")) + "") + "a") {
            case "hello":
                if false {
                    switch "" {
                    case "a":
                        v4 = (v4 / (((v4 ^ v4)) | 1))
                    case "hello":
                        v25 := GenS5{F0: 0}
                        _ = v25
                    }
                } else {
                    var v26 uint16 = 42
                    var v27 uint64 = 112
                    _ = v26
                    _ = v27
                }
            case "\n\t":
                for fi28 := 0; fi28 < 14; fi28++ {
                    v11 += v11
                    _ = fi28
                }
            case "":
                v10 = (v10 & int32(fi23))
            }
            _ = fi23
            _ = v24
        }
    } else {
        if len(v2) > 0 {
            v2[0] = 315
        }
    }
    if true {
        v4 = (v4 << ((42) & 63))
        v3 = 42
    } else {
        fmt.Println(v1, v10, v11)
        if true {
            var v29 float32 = (0.1 * v1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                fmt.Println(v3, v3)
                _ = len(v2)
                _ = fi30
            }
            _ = v29
        } else {
            for fi31 := 0; fi31 < 18; fi31++ {
                for fi32 := 0; fi32 < 11; fi32++ {
                    v33 := GenS5{F0: 296}
                    fmt.Println(v4, v1)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi32
                    _ = v33
                }
                _ = fi31
            }
            if false {
                if false {
                    var v34 int8 = genFn0(0, (v10 << ((4294967295) & 63)), v3)
                    v3 = int64(v10)
                    _ = v34
                }
                v10 = v10
            } else {
                _ = len(v2)
            }
        }
    }
    if (false || true) {
        for fi35 := 0; fi35 < 18; fi35++ {
            fmt.Println(v11, v3)
            fmt.Println(v11, fi35, v11)
            _ = fi35
        }
        switch "" {
        case "\n\t":
            var v36 int8 = v4
            _ = v36
        case "":
            v37 := GenS1{F0: false, F1: "foo bar", F2: (v11 / ((v11) | 1))}
            _ = v37
        case "a":
            v3 += (v3 ^ v3)
        default:
            var v38 uint16 = 32767
            _ = v38
        }
    } else {
        for fi39 := 0; fi39 < 18; fi39++ {
            if (true && true) {
                for fi40 := 0; fi40 < 2; fi40++ {
                    v11 -= v11
                    v4 += (v4 / ((100) | 1))
                    _ = fi40
                }
                v11 = (v11 & (^(v11)))
            } else {
                v1 = ((float32(0.0) * 1e10) * (v1 - 1e10))
            }
            for fi41 := 0; fi41 < 14; fi41++ {
                v11 += v11
                if ((-(v4)) == v4) {
                    v11 = 18446744073709551615
                    v11 -= v11
                }
                _ = fi41
            }
            _ = fi39
        }
        var v42 float64 = (1.5 + 0.0)
        v42 = 0.1
        _ = v42
    }
    var v43 uint = 409
    var v44 uint16 = 2
    switch -815 {
    case 592:
        v1 = v1
    default:
        var v45 uint8 = 0
        _ = v45
    }
    var v46 int8 = (v4 * 1)
    v47 := GenS5{F0: 2}
    v4 += 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v46
    _ = v47
}
