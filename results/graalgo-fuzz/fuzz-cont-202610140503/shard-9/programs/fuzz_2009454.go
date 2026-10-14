package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint64 }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]string)
    var v3 string = (((("\n\t" + ("foo bar" + "\n\t")) + ("hello" + "foo bar")) + (("\n\t" + ((((("foo bar" + ("foo bar" + "\n\t")) + "foo bar") + "foo bar") + "foo bar") + "a")) + "hello")) + "hello")
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 uint64 = 0
        fmt.Println(v5, v3)
        _ = fi4
        _ = v5
    }
    switch v3 {
    case "a":
        var v6 string = v3
        _ = v6
    case "foo bar":
        var v7 float64 = -2.25
        _ = v7
    default:
        switch 4294967295 {
        case 127:
            fmt.Println(v3)
        default:
            v8 := GenS1{F0: v3, F1: 126, F2: 127}
            _ = v8
        }
    }
    fmt.Println(v3, v3)
    var v9 uint8 = 1
    fmt.Println(v1, v1, v9)
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            var v11 uint8 = v9
            if false {
                for fi12 := 0; fi12 < 1; fi12++ {
                    fmt.Println(fi12, fi12, v3)
                    var v13 bool = false
                    _ = fi12
                    _ = v13
                }
                var v14 string = "\n\t"
                switch fi10 {
                case 4294967295:
                    var v15 string = ("\n\t" + v3)
                    _ = v15
                case 128:
                    fmt.Println(fi10, fi10, v9)
                case 245:
                    fmt.Println(v11, v11, v1)
                default:
                    if (v9 == v9) {
                        for fi16 := 0; fi16 < 15; fi16++ {
                            if ((v11 <= v9) || false) {
                                var v17 int8 = 92
                                _ = v17
                            } else {
                                var v18 uint64 = 68
                                var v19 uint32 = 4294967295
                                v20 := []int{42, 425, fi10}
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            _ = fi16
                        }
                    }
                }
                _ = v14
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                v22 := []int{(fi10 & -2147483648), fi10}
                if (-1 >= 1) {
                    if (-1 <= -186) {
                        var v23 int8 = int8(fi10)
                        _ = v23
                    }
                    v9 -= (v11 ^ v9)
                    v2[""] = ""
                } else {
                    var v24 bool = true
                    v24 = true
                    _ = v24
                }
                _ = fi21
                _ = v22
            }
            _ = fi10
            _ = v11
        }
        v25 := GenS1{F0: (v3 + v3), F1: 0, F2: 1}
        v26 := GenS2{F0: 2147483647, F1: 127, F2: 0}
        _ = v25
        _ = v26
    } else {
        v27 := GenS1{F0: v3, F1: 1, F2: v9}
        _ = v27
    }
    var v28 string = "a"
    if true {
        v1 += v1
        if true {
            v2[(("a" + v3) + "\n\t")] = "a"
        } else {
            v29 := v2[(v3 + (("a" + (("" + v28) + v28)) + (v3 + "a")))]
            fmt.Println(v3, v3)
            switch "hello" {
            case "foo bar":
                if true {
                    v30 := GenS1{F0: v29, F1: -1, F2: (v9 << (((v1 | 1)) & 63))}
                    v31 := v2[(v3 + "hello")]
                    var v32 uint8 = v9
                    _ = v30
                    _ = v31
                    _ = v32
                } else {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        if ((1e10 + 1.5) != -2.25) {
                            if (-1 >= 32767) {
                                v34 := GenS2{F0: 614, F1: int8(v9), F2: 775}
                                v28 = (("a" + "") + (("hello" + (v28 + ("hello" + "hello"))) + (v3 + v3)))
                                _ = v34
                            }
                            v35 := []int{}
                            _ = v35
                        } else {
                            v3 = "a"
                            for fi36 := 0; fi36 < 18; fi36++ {
                                fmt.Println(v28)
                                v37 := []int{0}
                                _ = fi36
                                _ = v37
                            }
                            v3 = v28
                        }
                        _ = fi33
                    }
                }
            case "a":
                v38 := []int{}
                _ = v38
            case "\n\t":
                v39 := GenS2{F0: 2147483648, F1: int8(v1), F2: 1}
                _ = v39
            }
            _ = v29
        }
    }
    fmt.Println(v1, v28, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v9
}
