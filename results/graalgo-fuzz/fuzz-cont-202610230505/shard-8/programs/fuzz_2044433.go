package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int32 = -579
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    v4 := GenS1{F0: 2147483647, F1: 170}
    for fi5 := 0; fi5 < 20; fi5++ {
        v1 = (v1 + (v1 | (v1 & 42)))
        v6 := []int{fi5}
        _ = fi5
        _ = v6
    }
    v7 := v2["\n\t"]
    switch 32767 {
    case 9223372036854775807:
        var v8 uint64 = 2
        _ = v8
    }
    if ("a" != "a") {
        if true {
            switch -922 {
            case -943:
                if (42 < 65535) {
                    v9 := v2["a"]
                    var v10 uint8 = 2
                    v1 -= v1
                    _ = v10
                    _ = v9
                } else {
                    v11 := GenS1{F0: 4294967295, F1: 70}
                    var v12 bool = (true || true)
                    _ = v11
                    _ = v12
                }
            default:
                v7 = (v7 + v7)
            }
            var v13 uint8 = 42
            _ = v13
        } else {
            v14 := []int{983}
            for fi15 := 0; fi15 < 12; fi15++ {
                var v16 int32 = v1
                for fi17 := 0; fi17 < 9; fi17++ {
                    if ("\n\t" < "foo bar") {
                        var v18 bool = (false && (false && (1e10 >= 3.14159)))
                        _ = v18
                    }
                    if false {
                        fmt.Println(v1, fi17, v1)
                        var v19 uint32 = 811
                        var v20 float32 = float32(1e10)
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int64 = 9223372036854775807
                        _ = v21
                    }
                    if false {
                        v22 := GenS1{F0: 1, F1: 128}
                        v23 := GenS1{F0: 4294967295, F1: 195}
                        v7 = ((v7 + "hello") + "")
                        _ = v22
                        _ = v23
                    }
                    _ = fi17
                }
                _ = fi15
                _ = v16
            }
            if (-1 != 0) {
                for fi24 := 0; fi24 < 16; fi24++ {
                    switch (("a" + v7) + "") {
                    case "foo bar":
                        if (true && (((1e10 * (0.0 / float32(1e10))) - 0.1) != 3.14159)) {
                            var v25 int32 = v1
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 14; fi26++ {
                                var v27 int8 = int8(fi24)
                                fmt.Println(fi26, v7, v7)
                                _ = fi26
                                _ = v27
                            }
                        }
                    }
                    if false {
                        var v28 int64 = -618
                        _ = v28
                    } else {
                        var v29 int8 = -1
                        var v30 int8 = v29
                        if (true || ((!(true || (true && false))) || true)) {
                            if true {
                                fmt.Println(v1, v7, v7)
                                v7 = "foo bar"
                            }
                            var v31 bool = false
                            var v32 float32 = 1e10
                            _ = v31
                            _ = v32
                        }
                        _ = v29
                        _ = v30
                    }
                    delete(v3, v7)
                    _ = fi24
                }
                v1 += (-(v1))
                v33 := GenS1{F0: 65535, F1: 255}
                _ = v33
            }
            _ = v14
        }
    } else {
        fmt.Println(v1)
        var v34 int16 = 1000
        for fi35 := 0; fi35 < 12; fi35++ {
            var v36 float32 = 0.1
            for fi37 := 0; fi37 < 10; fi37++ {
                var v38 int8 = 42
                _ = fi37
                _ = v38
            }
            for fi39 := 0; fi39 < 4; fi39++ {
                v40 := GenS1{F0: 856, F1: 255}
                switch (fi35 % ((-303) | 1)) {
                case -2147483648:
                    v34 += v34
                default:
                    fmt.Println(v7, v7)
                }
                _ = fi39
                _ = v40
            }
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    var v41 float64 = -2.25
    fmt.Println(v41, v7, v41)
    var v42 uint32 = 2147483648
    v43 := GenS1{F0: v42, F1: 127}
    v44 := v2["a"]
    var v45 int64 = -960
    fmt.Println(v1, v44)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v7
}
