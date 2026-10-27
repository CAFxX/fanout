package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 string }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint }
func genFn0() int8 {
    var v6 int8 = -128
    var v7 uint8 = 125
    v8 := GenS1{F0: 507, F1: 2, F2: ""}
    _ = v6
    _ = v7
    _ = v8
    return v6
}
func genFn1(p0 int, p1 int64, p2 uint16) int16 {
    fmt.Println(p0)
    var v9 float32 = 1e10
    var v10 uint = 513
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = -72
    v3 := []string{"", v1, v1, v1}
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    switch 0 {
    case -9223372036854775808:
        v1 = ""
    case 32768:
        fmt.Println(v1, v2)
    }
    switch 0 {
    case 0:
        if false {
            v11 := GenS1{F0: 2, F1: 1, F2: (v1 + (v1 + v1))}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 int = 4294967295
                for fi14 := 0; fi14 < 6; fi14++ {
                    v15 := GenS2{F0: int32(v2), F1: 42}
                    var v16 int64 = 0
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                v1 = "foo bar"
                _ = fi12
                _ = v13
            }
            v17 := v5[42]
            _ = v17
        }
    }
    fmt.Println(v1, v2)
    if true {
        if (4294967295 != 56) {
            var v18 bool = true
            switch 812 {
            case -1:
                for fi19 := 0; fi19 < 2; fi19++ {
                    if true {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            var v21 uint32 = 456
                            _ = fi20
                            _ = v21
                        }
                        switch fi19 {
                        case 1:
                            v18 = (true || true)
                        }
                        if ((1e10 + -2.25) != 3.14159) {
                            var v22 float64 = 1e10
                            var v23 string = v1
                            _ = v22
                            _ = v23
                        } else {
                            var v24 int32 = -546
                            _ = v24
                        }
                    }
                    for fi25 := 0; fi25 < 1; fi25++ {
                        switch (fi25 | (fi19 | 42)) {
                        case 127:
                            v26 := GenS3{F0: genFn1(fi19, -1, 65535), F1: 2147483647}
                            _ = v26
                        }
                        _ = fi25
                    }
                    if (!(float32(1e10) < -2.25)) {
                        for fi27 := 0; fi27 < 11; fi27++ {
                            fmt.Println(fi19, fi19, v1)
                            fmt.Println(v1, v18, v18)
                            _ = fi27
                        }
                    }
                    _ = fi19
                }
            }
            _ = v18
        }
        for fi28 := 0; fi28 < 8; fi28++ {
            fmt.Println(fi28)
            if false {
                if ((650 <= -224) && ("\n\t" < v1)) {
                    if true {
                        v4[1] = false
                        var v29 int = (fi28 * (fi28 * (fi28 * fi28)))
                        _ = v29
                    } else {
                        var v30 float64 = (1.5 - 0.1)
                        v31 := GenS1{F0: 42, F1: 42, F2: v1}
                        fmt.Println(fi28)
                        _ = v30
                        _ = v31
                    }
                }
                var v32 int8 = -76
                for fi33 := 0; fi33 < 2; fi33++ {
                    var v34 int64 = -1
                    v34 = (v34 & (v34 | 760))
                    _ = fi33
                    _ = v34
                }
                _ = v32
            } else {
                v2 -= v2
                fmt.Println(v2, fi28)
                var v35 int16 = 1
                _ = v35
            }
            v2 -= v2
            _ = fi28
        }
    }
    delete(v5, 296)
    var v36 int64 = -740
    var v37 uint32 = 380
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v38 := GenS2{F0: int32(v36), F1: (v36 + (v36 ^ -919))}
    v37 = 0
    var v39 string = "foo bar"
    v40 := GenS4{F0: 0}
    if (false && true) {
        for fi41 := 0; fi41 < 8; fi41++ {
            var v42 bool = true
            _ = fi41
            _ = v42
        }
        var v43 float32 = (1.5 / (float32(((1e10 * float64(3.14159)) + ((0.1 - 0.0) / 1.0))) * 1e10))
        v37 += 65535
        _ = v43
    } else {
        v37 += v37
        switch v1 {
        case "\n\t":
            v37 += (v37 >> ((153) & 63))
        default:
            v44 := v4[v37]
            _ = v44
        }
    }
    var v45 uint8 = 42
    v36 = v36
    _ = len(v3)
    var v46 string = v39
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v46
    _ = v5
}
