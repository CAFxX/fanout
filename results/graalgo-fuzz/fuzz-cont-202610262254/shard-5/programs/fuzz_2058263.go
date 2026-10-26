package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8, p1 uint64) string {
    var v7 float64 = (3.14159 * 0.1)
    if true {
        fmt.Println(p0, p1)
        v8 := []int{9223372036854775807, 0, -447, 42}
        _ = v8
    } else {
        v9 := []int{42, 0}
        fmt.Println(v7, p1, p1)
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(p1, fi10)
        _ = fi10
    }
    fmt.Println(v7, p0)
    _ = p0
    _ = p1
    _ = v7
    return "a"
}

func main() {
    var v1 bool = false
    v2 := []string{"a", "foo bar", "", "foo bar"}
    v3 := []string{"a", (("" + "") + ("" + "a"))}
    v4 := make(map[int]string)
    var v5 uint32 = 1
    var v6 int64 = 42
    var v11 uint16 = 1
    v12 := GenS1{F0: v5, F1: 128}
    if v1 {
        v1 = true
        if v1 {
            if v1 {
                v13 := v4[32768]
                v1 = v1
                switch 128 {
                case 4294967295:
                    v5 -= (v5 & v5)
                case 32768:
                    v14 := GenS1{F0: v5, F1: 128}
                    _ = v14
                case 42:
                    fmt.Println(v1, v1, v11)
                default:
                    v5 = 2147483648
                }
                _ = v13
            } else {
                switch "\n\t" {
                case "hello":
                    var v15 int16 = -984
                    _ = v15
                }
            }
            v1 = (((!false) && v1) && (!(v1 || false)))
        }
    }
    var v16 int8 = 0
    v17 := v4[216]
    if (v17 <= "a") {
        if true {
            v18 := GenS1{F0: v5, F1: 2}
            _ = v18
        } else {
            var v19 int64 = v6
            _ = v19
        }
        v6 = v6
        switch v17 {
        case "foo bar":
            var v20 uint32 = v5
            _ = v20
        }
    } else {
        switch 42 {
        case 1:
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := GenS2{F0: 42}
                switch fi21 {
                case 128:
                    v3 = append(v3, "a")
                case -1:
                    v6 += (v6 ^ v6)
                case 1:
                    for fi23 := 0; fi23 < 14; fi23++ {
                        var v24 uint = 448
                        _ = fi23
                        _ = v24
                    }
                }
                var v25 uint32 = v5
                _ = fi21
                _ = v22
                _ = v25
            }
        case 2147483647:
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 float64 = 1.5
                v5 = v5
                _ = fi26
                _ = v27
            }
        case -9223372036854775808:
            if (1e10 <= ((float64((1e10 / 1.5)) / ((((0.1 + -2.25) / 1.0) - 0.0) * 3.14159)) + (3.14159 * 3.14159))) {
                v11 -= (v11 - 32767)
                var v28 string = v17
                _ = v28
            } else {
                v5 = (v5 & v5)
            }
        }
        if ("" != "\n\t") {
            var v29 float64 = 1.5
            _ = v29
        }
        var v30 uint16 = (v11 + (^(v11)))
        _ = v30
    }
    switch "" {
    case "hello":
        if (2 == 2) {
            v17 = (v17 + (v17 + v17))
        } else {
            var v31 int = 1
            var v32 uint32 = v5
            v32 -= uint32(v11)
            _ = v31
            _ = v32
        }
    }
    for fi33 := 0; fi33 < 4; fi33++ {
        v34 := GenS1{F0: v5, F1: 255}
        v35 := GenS3{F0: (true && (v11 > 1))}
        switch 9223372036854775807 {
        case 32767:
            _ = len(v4)
        case 127:
            var v36 uint = 0
            _ = v36
        case 42:
            if (!v1) {
                v2 = append(v2, (v17 + ((("\n\t" + ("\n\t" + v17)) + "hello") + v17)))
                for fi37 := 0; fi37 < 5; fi37++ {
                    switch v17 {
                    case "foo bar":
                        if (1.5 > float64(1.5)) {
                            switch v17 {
                            case "":
                                if ((^(v5)) == (v5 / ((v5) | 1))) {
                                    var v38 bool = ((560 != 828) && false)
                                    _ = v38
                                }
                            case "foo bar":
                                if v1 {
                                    var v39 float32 = 1e10
                                    _ = v39
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            default:
                                v11 -= 2
                            }
                            fmt.Println(v11, v6)
                            fmt.Println(v11, v5, v17)
                        } else {
                            v17 = ((v17 + "\n\t") + (v17 + ("a" + (("" + "a") + v17))))
                        }
                    case "":
                        _ = len(v4)
                    default:
                        if ((^(v16)) <= v16) {
                            fmt.Println(v11, v11, v6)
                            v1 = (!((v16 + v16) < v16))
                        }
                    }
                    _ = fi37
                }
                switch (((("\n\t" + "\n\t") + "foo bar") + (v17 + ("\n\t" + ("a" + (v17 + ("\n\t" + "a")))))) + v17) {
                case "hello":
                    _ = len(v3)
                default:
                    if false {
                        switch (-(fi33)) {
                        case -9223372036854775808:
                            delete(v4, -9223372036854775808)
                        case 128:
                            switch v17 {
                            case "a":
                                fmt.Println(v16)
                            default:
                                var v40 int32 = 0
                                _ = v40
                            }
                        case -2147483648:
                            v6 = 1
                        }
                    }
                }
            } else {
                var v41 uint = 255
                var v42 float32 = 1e10
                _ = v41
                _ = v42
            }
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    if false {
        v2 = append(v2, v17)
    }
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 float64 = ((0.1 - (0.1 - (((1e10 / (3.14159 - 1.5)) * 0.1) - ((float64(-2.25) * 1e10) - 3.14159)))) - 1.5)
        _ = fi43
        _ = v44
    }
    var v45 int16 = -1
    var v46 int = 757
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
