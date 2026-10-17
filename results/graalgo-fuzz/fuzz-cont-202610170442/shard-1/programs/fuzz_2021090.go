package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
func genFn0() int32 {
    var v4 bool = (false || false)
    if (((!(!false)) && v4) || v4) {
        v5 := GenS1{F0: 875}
        var v6 int = 4294967295
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 42}
    fmt.Println(v4)
    _ = v4
    _ = v7
    return -2147483648
}
func genFn1(p0 uint8) uint {
    var v8 uint32 = 362
    v9 := GenS2{F0: 255}
    _ = p0
    _ = v8
    _ = v9
    return 362
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := []string{"a"}
    switch 0 {
    case -277:
        if false {
            var v10 int8 = 0
            var v11 float64 = ((1e10 * (3.14159 - 1.5)) + float64((3.14159 / -2.25)))
            _ = v10
            _ = v11
        }
    case 128:
        var v12 uint64 = 0
        _ = v12
    case 32767:
        var v13 int16 = -32768
        _ = v13
    }
    var v14 uint64 = 203
    switch 128 {
    case 0:
        for fi15 := 0; fi15 < 12; fi15++ {
            v14 += 0
            var v16 int = fi15
            v16 = (v16 ^ v16)
            _ = fi15
            _ = v16
        }
    case -1:
        v17 := GenS2{F0: 2}
        _ = v17
    case 4294967295:
        var v18 uint = 2
        _ = v18
    }
    var v19 uint64 = v14
    var v20 int64 = 199
    var v21 uint32 = 65535
    v22 := GenS1{F0: -1}
    switch -356 {
    case 128:
        var v23 float64 = (-2.25 + 1.5)
        _ = v23
    case -297:
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS1{F0: -2147483648}
            if false {
                var v26 uint16 = 42
                var v27 int64 = (v20 >> ((0) & 63))
                for fi28 := 0; fi28 < 5; fi28++ {
                    fmt.Println(v14, v14)
                    var v29 int = fi28
                    _ = fi28
                    _ = v29
                }
                _ = v26
                _ = v27
            }
            _ = fi24
            _ = v25
        }
    }
    switch -2147483648 {
    case -2147483648:
        switch -755 {
        case -814:
            var v30 uint = 327
            _ = v30
        case -9223372036854775808:
            if (v20 == v20) {
                switch 9223372036854775807 {
                case -1:
                    switch 429 {
                    case -1:
                        _ = len(v3)
                    default:
                        _ = len(v1)
                    }
                case 42:
                    var v31 uint64 = v14
                    _ = v31
                case 32768:
                    var v32 uint = 1
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 8; fi33++ {
                        switch fi33 {
                        case 127:
                            v34 := GenS1{F0: -664}
                            _ = v34
                        case 985:
                            switch fi33 {
                            case 904:
                                _ = len(v2)
                            case 9223372036854775807:
                                v35 := GenS1{F0: -342}
                                _ = v35
                            }
                        case 502:
                            var v36 string = "hello"
                            _ = v36
                        default:
                            delete(v2, v21)
                        }
                        switch (("\n\t" + "\n\t") + ("" + "a")) {
                        case "\n\t":
                            var v37 float32 = ((0.1 / float32(1e10)) - (0.0 * float32((((0.1 * float64(1e10)) - 1.5) * -2.25))))
                            _ = v37
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi33
                    }
                }
            }
        }
    case -90:
        if (!(uint(v14) <= genFn1(42))) {
            v38 := GenS1{F0: 2147483647}
            if false {
                if false {
                    v19 -= 42
                    v19 = v19
                }
                for fi39 := 0; fi39 < 20; fi39++ {
                    var v40 int8 = 111
                    if ((!(!true)) && true) {
                        v19 = (v14 | v19)
                        v41 := GenS3{F0: true}
                        _ = v41
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi39
                    _ = v40
                }
            } else {
                _ = len(v3)
                var v42 float32 = (0.1 + 3.14159)
                _ = v42
            }
            var v43 float32 = float32((-2.25 / 1.0))
            _ = v38
            _ = v43
        }
    }
    v44 := GenS1{F0: 1}
    for fi45 := 0; fi45 < 1; fi45++ {
        v14 -= 4294967295
        v46 := GenS4{F0: true}
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi45
        _ = v46
    }
    var v47 int8 = 127
    v47 = int8(v21)
    for fi48 := 0; fi48 < 12; fi48++ {
        v20 -= (-(v20))
        v19 -= 2
        _ = fi48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v44
    _ = v47
}
