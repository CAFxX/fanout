package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int32 }
type GenS2 struct { F0 string; F1 bool; F2 int32 }

func main() {
    var v1 int = 127
    v2 := make(map[string]bool)
    var v3 int32 = 1
    v4 := []string{("foo bar" + "\n\t"), "\n\t", ((("hello" + "") + "\n\t") + ("a" + "\n\t"))}
    v3 += v3
    for fi5 := 0; fi5 < 20; fi5++ {
        v1 += (fi5 >> ((681) & 63))
        fmt.Println(v1, v1, v1)
        _ = fi5
    }
    v6 := GenS1{F0: -128, F1: true, F2: v3}
    _ = len(v4)
    if ((v3 + 385) > v3) {
        v3 += (v3 * (-(v3)))
    } else {
        if false {
            if (true || (255 < 537)) {
                if false {
                    fmt.Println(v1, v3)
                    fmt.Println(v1, v3)
                }
                if ((0 == 65535) && false) {
                    v7 := v2[""]
                    v1 -= 892
                    var v8 uint64 = 2
                    _ = v7
                    _ = v8
                } else {
                    var v9 uint32 = 488
                    for fi10 := 0; fi10 < 15; fi10++ {
                        v2["\n\t"] = true
                        if false {
                            v3 += (v3 - -2147483648)
                            v4 = append(v4, "hello")
                            fmt.Println(v9, v3)
                        }
                        for fi11 := 0; fi11 < 3; fi11++ {
                            fmt.Println(v1, fi11)
                            if true {
                                var v12 int8 = 1
                                v13 := GenS1{F0: 100, F1: (65535 != 994), F2: v3}
                                _ = v12
                                _ = v13
                            } else {
                                v14 := v2["\n\t"]
                                v1 = fi10
                                if false {
                                    if (false || false) {
                                        if ((3.14159 / (1e10 - -2.25)) >= (1.5 * float32(1.5))) {
                                            _ = len(v2)
                                            v14 = (uint64(v9) != 2)
                                            if ((-2.25 / 0.1) < float64(float32(1.5))) {
                                                v3 = (v3 + v3)
                                                v15 := GenS1{F0: -84, F1: v14, F2: v3}
                                                _ = len(v2)
                                                _ = v15
                                            }
                                        }
                                        if (!((v3 < (v3 << ((890) & 63))) && v14)) {
                                            if v14 {
                                                v9 += (v9 - (v9 | v9))
                                                fmt.Println(v9, v9)
                                                delete(v2, "foo bar")
                                            } else {
                                                _ = len(v4)
                                                var v16 bool = v14
                                                _ = v16
                                            }
                                            fmt.Println(v3, fi11)
                                            var v17 int64 = 0
                                            fmt.Println(v3, v1, v9)
                                            v18 := GenS2{F0: "", F1: true, F2: -2147483648}
                                            _ = v17
                                            _ = v18
                                        }
                                    } else {
                                        var v19 int = (v1 & (v1 - int(v3)))
                                        if (!false) {
                                            var v20 uint32 = (^(v9))
                                            var v21 float64 = ((1.5 * ((0.0 / 1.0) + (1.5 * 3.14159))) - 1e10)
                                            _ = v20
                                            _ = v21
                                        }
                                        v14 = (!(-9223372036854775808 >= 0))
                                        _ = v19
                                    }
                                    if v14 {
                                        v22 := GenS2{F0: "hello", F1: v14, F2: v3}
                                        fmt.Println(fi11, v3)
                                        var v23 float64 = (0.0 / (1e10 + 1.5))
                                        v3 = (v3 | v3)
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                _ = v14
                            }
                            switch "hello" {
                            case "foo bar":
                                var v24 uint8 = 0
                                _ = v24
                            }
                            _ = fi11
                        }
                        _ = fi10
                    }
                    _ = v9
                }
            }
        }
        if true {
            v25 := v2["hello"]
            var v26 bool = ((0 > 127) && v25)
            v4 = append(v4, ("foo bar" + "hello"))
            _ = v25
            _ = v26
        } else {
            v1 = 2147483647
            var v27 uint64 = 250
            var v28 float32 = 3.14159
            _ = v27
            _ = v28
        }
        v3 -= 0
    }
    if true {
        v3 = v3
        fmt.Println(v1)
    }
    var v29 uint64 = 4294967295
    var v30 uint64 = (v29 << ((255) & 63))
    fmt.Println(v1, v1)
    fmt.Println(v30)
    v3 -= (v3 + 649)
    if true {
        var v31 float32 = 1e10
        _ = v31
    }
    v32 := GenS2{F0: ("hello" + "foo bar"), F1: true, F2: (v3 * v3)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v6
}
