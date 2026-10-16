package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int16) int64 {
    p0 = -291
    v7 := []int{127, 868, 62, 128}
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 1
    var v4 int32 = 2147483647
    v5 := make(map[uint32]bool)
    v6 := make(map[string]uint32)
    v8 := GenS1{F0: 569}
    fmt.Println(v1, v4, v1)
    switch -982 {
    case 9223372036854775807:
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint8 = 2
            _ = fi9
            _ = v10
        }
    default:
        fmt.Println(v1, v3)
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if (72 != v3) {
            if (1000 <= 0) {
                v12 := []int{(fi11 << ((2) & 63)), fi11}
                _ = v12
            } else {
                switch fi11 {
                case 1:
                    fmt.Println(fi11, v3)
                case 0:
                    if true {
                        fmt.Println(v4, v1)
                    }
                }
                var v13 float32 = 0.1
                _ = v13
            }
            v14 := []int{(fi11 - fi11)}
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 uint32 = 1
                fmt.Println(fi15)
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        if (-113 >= -5) {
            if (((!false) && (!(("\n\t" + "") > "\n\t"))) && true) {
                switch (fi11 / ((fi11) | 1)) {
                case -2147483648:
                    if (9223372036854775807 > 42) {
                        switch "foo bar" {
                        case "hello":
                            switch "\n\t" {
                            case "":
                                fmt.Println(v4, v3, v3)
                            case "foo bar":
                                fmt.Println(fi11, v3)
                            case "a":
                                fmt.Println(v3, v1)
                            }
                        }
                        switch (fi11 >> ((18446744073709551615) & 63)) {
                        case -2147483648:
                            var v17 uint32 = 4294967295
                            _ = v17
                        case 128:
                            v1 = v1
                        case -991:
                            var v18 uint = 985
                            _ = v18
                        default:
                            v19 := []int{-9223372036854775808, (-(fi11)), fi11, fi11}
                            _ = v19
                        }
                        v20 := []int{(fi11 - fi11)}
                        _ = v20
                    } else {
                        fmt.Println(v3, v1)
                        switch fi11 {
                        case -1:
                            for fi21 := 0; fi21 < 17; fi21++ {
                                fmt.Println(fi21, v1, v4)
                                var v22 uint64 = v1
                                v3 = 42
                                _ = fi21
                                _ = v22
                            }
                        default:
                            switch "\n\t" {
                            case "":
                                v23 := GenS1{F0: 42}
                                _ = v23
                            }
                        }
                        var v24 uint = 255
                        _ = v24
                    }
                case 741:
                    var v25 float32 = float32(3.14159)
                    _ = v25
                case 966:
                    var v26 float64 = 0.0
                    _ = v26
                }
            } else {
                if true {
                    fmt.Println(v1)
                    for fi27 := 0; fi27 < 16; fi27++ {
                        var v28 float64 = 3.14159
                        switch ("a" + "hello") {
                        case "hello":
                            var v29 uint64 = (v1 | 1)
                            _ = v29
                        case "a":
                            var v30 uint = 414
                            _ = v30
                        case "":
                            switch fi11 {
                            case 4294967295:
                                var v31 uint64 = v1
                                _ = v31
                            case -1:
                                v4 -= (v4 << ((0) & 63))
                            }
                        }
                        _ = fi27
                        _ = v28
                    }
                } else {
                    fmt.Println(v1)
                    var v32 uint32 = 0
                    _ = v32
                }
                fmt.Println(v3, v4)
            }
        } else {
            delete(v5, 2147483648)
        }
        var v33 bool = true
        _ = fi11
        _ = v33
    }
    if (true && true) {
        var v34 uint = 2
        var v35 uint16 = 65535
        _ = v34
        _ = v35
    } else {
        v36 := GenS1{F0: 0}
        v3 = (v3 - (v3 + v3))
        _ = v36
    }
    switch -2147483648 {
    case 9223372036854775807:
        var v37 uint16 = 32767
        _ = v37
    default:
        v38 := GenS1{F0: 32767}
        _ = v38
    }
    v39 := []int{-1, 32768}
    v6[(("hello" + ("\n\t" + "a")) + ("\n\t" + ("foo bar" + "\n\t")))] = 2
    if ((!true) || true) {
        var v40 uint32 = 524
        var v41 int64 = -1
        _ = v40
        _ = v41
    }
    var v42 uint16 = 0
    for fi43 := 0; fi43 < 14; fi43++ {
        var v44 float64 = 0.1
        var v45 float32 = (1e10 + -2.25)
        _ = fi43
        _ = v44
        _ = v45
    }
    var v46 uint = 124
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v46
    _ = v5
    _ = v6
    _ = v8
}
