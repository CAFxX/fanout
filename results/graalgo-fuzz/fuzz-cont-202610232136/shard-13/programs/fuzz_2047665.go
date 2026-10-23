package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint }
type GenS3 struct { F0 uint16; F1 int8 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 string = ""
    var v3 bool = (!((v1 < (v1 | 42)) || true))
    var v4 float32 = 1.5
    v5 := make(map[string]bool)
    if ((!v3) && false) {
        switch (v2 + v2) {
        case "hello":
            for fi6 := 0; fi6 < 9; fi6++ {
                for fi7 := 0; fi7 < 12; fi7++ {
                    v8 := GenS1{F0: 100, F1: 65535}
                    v9 := GenS2{F0: 42, F1: 2147483647, F2: 170}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                if v3 {
                    _ = len(v5)
                    v10 := v5[(v2 + v2)]
                    for fi11 := 0; fi11 < 5; fi11++ {
                        var v12 int = (fi11 | (-(fi11)))
                        switch v2 {
                        case "hello":
                            switch "foo bar" {
                            case "hello":
                                v3 = (!(!true))
                            case "a":
                                if true {
                                    v13 := []int{}
                                    _ = v13
                                }
                            default:
                                if true {
                                    var v14 uint32 = 768
                                    v15 := []int{(^(v12)), (^(fi11))}
                                    _ = v14
                                    _ = v15
                                }
                            }
                        case "a":
                            if v10 {
                                v16 := GenS2{F0: -1, F1: 0, F2: 0}
                                v12 -= (v12 | -1)
                                _ = v16
                            } else {
                                v17 := []int{0, v12, (fi6 - fi6), 1}
                                v4 += 0.1
                                fmt.Println(v4, v10, fi6)
                                _ = v17
                            }
                        case "foo bar":
                            v10 = v10
                            v10 = v3
                            v2 = v2
                        }
                        v2 = "\n\t"
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                }
                _ = fi6
            }
        }
    }
    var v18 uint16 = 255
    var v19 uint32 = 1
    v1 -= -58
    if (((false && ((v4 - v4) < float32(1e10))) && ((true || v3) || (!v3))) && (0 >= 1)) {
        var v20 int16 = 42
        _ = v20
    } else {
        if v3 {
            var v21 int64 = (v1 ^ (v1 % (((v1 ^ v1)) | 1)))
            _ = v21
        } else {
            v22 := v5[(v2 + ((v2 + "\n\t") + ("" + ("foo bar" + "\n\t"))))]
            _ = v22
        }
    }
    v23 := GenS2{F0: 42, F1: 65535, F2: 1}
    fmt.Println(v3)
    if (1 <= 0) {
        var v24 int32 = -1
        v2 = v2
        if (v3 || v3) {
            switch 72 {
            case 2147483647:
                for fi25 := 0; fi25 < 16; fi25++ {
                    switch v2 {
                    case "":
                        var v26 string = "hello"
                        _ = v26
                    case "\n\t":
                        for fi27 := 0; fi27 < 6; fi27++ {
                            v28 := []int{}
                            var v29 int16 = -1
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = fi25
                }
            case 4294967295:
                _ = len(v5)
            }
            if false {
                var v30 float64 = 3.14159
                _ = v30
            }
            var v31 uint32 = (v19 >> ((0) & 63))
            _ = v31
        }
        _ = v24
    }
    v32 := GenS3{F0: v18, F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
