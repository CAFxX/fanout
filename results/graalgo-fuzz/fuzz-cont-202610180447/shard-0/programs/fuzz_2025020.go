package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint16) uint8 {
    switch 4294967295 {
    case 175:
        fmt.Println(p0)
    case -616:
        switch 379 {
        case 247:
            if (-1 == 331) {
                v6 := GenS1{F0: 127}
                var v7 int16 = 32767
                v8 := GenS1{F0: 255}
                _ = v6
                _ = v7
                _ = v8
            } else {
                v9 := GenS1{F0: 128}
                _ = v9
            }
        case 42:
            v10 := []int{127, 1, 0, -73}
            _ = v10
        }
    default:
        p0 = p0
    }
    v11 := GenS1{F0: 128}
    _ = p0
    _ = v11
    return 183
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    var v4 int = -9223372036854775808
    v5 := make(map[int]bool)
    if true {
        fmt.Println(v4)
        v4 += v4
    }
    var v12 string = "\n\t"
    var v13 int = (v4 << ((255) & 63))
    v14 := GenS1{F0: 0}
    if true {
        var v15 int8 = -128
        if true {
            var v16 uint8 = 255
            _ = v16
        }
        v17 := GenS2{F0: 4294967295, F1: -1}
        _ = v15
        _ = v17
    }
    fmt.Println(v13, v13)
    v4 -= v13
    v18 := GenS1{F0: 161}
    switch v12 {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := GenS1{F0: 128}
                _ = fi19
                _ = v20
            }
        case "hello":
            for fi21 := 0; fi21 < 1; fi21++ {
                if false {
                    var v22 int = fi21
                    _ = v22
                }
                v23 := v2["a"]
                _ = fi21
                _ = v23
            }
        case "a":
            var v24 int16 = 32767
            _ = v24
        default:
            var v25 int16 = -32768
            _ = v25
        }
    case "\n\t":
        if true {
            if (!(42 > -1)) {
                switch (v13 | (v13 * (v4 >> ((925) & 63)))) {
                case -616:
                    var v26 uint8 = genFn0(533)
                    _ = v26
                case 9223372036854775807:
                    var v27 int32 = 42
                    _ = v27
                case 535:
                    v28 := GenS3{F0: 4294967295}
                    _ = v28
                }
            }
            switch v12 {
            case "a":
                var v29 int8 = -128
                _ = v29
            default:
                v1[("a" + "")] = 42
            }
            var v30 uint32 = 2147483648
            _ = v30
        } else {
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 int = -2147483648
                var v33 int16 = 42
                _ = fi31
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 18; fi34++ {
                for fi35 := 0; fi35 < 8; fi35++ {
                    var v36 int8 = 100
                    switch -323 {
                    case -9223372036854775808:
                        switch v13 {
                        case 42:
                            if true {
                                switch (((v12 + "hello") + v12) + "a") {
                                case "a":
                                    var v37 uint16 = 1
                                    _ = v37
                                case "\n\t":
                                    var v38 string = (v12 + "hello")
                                    _ = v38
                                default:
                                    v4 -= v4
                                }
                                v39 := GenS3{F0: 42}
                                _ = v39
                            }
                        default:
                            var v40 int64 = -1
                            _ = v40
                        }
                    case 32767:
                        v41 := GenS1{F0: 255}
                        _ = v41
                    }
                    _ = fi35
                    _ = v36
                }
                _ = fi34
            }
            var v42 int16 = -32768
            _ = v42
        }
    case "hello":
        var v43 bool = true
        _ = v43
    }
    v12 = v12
    var v44 bool = (!false)
    var v45 int32 = 1
    v4 -= v13
    v12 = v12
    v4 = -1
    v46 := GenS4{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
