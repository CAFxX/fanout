package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int64; F1 int16; F2 int64 }

func main() {
    var v1 int8 = 103
    v2 := make(map[uint32]uint32)
    var v3 int32 = -369
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v6 := make(map[string]string)
    var v7 bool = true
    if v7 {
        if (212 > 0) {
            var v8 uint = 18446744073709551615
            var v9 string = ((("a" + "a") + "") + "\n\t")
            _ = v8
            _ = v9
        } else {
            v3 -= v3
            fmt.Println(v3, v1, v7)
        }
    }
    v10 := []int{32767}
    v3 = v3
    fmt.Println(v3, v7)
    delete(v5, 2147483647)
    switch ("foo bar" + "hello") {
    case "\n\t":
        var v11 bool = true
        _ = v11
    default:
        v12 := v5[-616]
        _ = v12
    }
    fmt.Println(v1)
    var v13 uint = 42
    var v14 uint16 = 0
    if (213 >= 2) {
        var v15 int = 127
        _ = v15
    } else {
        v3 = 0
        for fi16 := 0; fi16 < 17; fi16++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            var v17 string = ((("\n\t" + "a") + "\n\t") + "\n\t")
            _ = fi16
            _ = v17
        }
        if v7 {
            if (!("\n\t" < "")) {
                for fi18 := 0; fi18 < 5; fi18++ {
                    v7 = (v3 < 42)
                    if (!(v7 || false)) {
                        v2[2147483647] = 1
                    }
                    _ = fi18
                }
            } else {
                if v7 {
                    for fi19 := 0; fi19 < 7; fi19++ {
                        var v20 float32 = 1.5
                        switch fi19 {
                        case -577:
                            for fi21 := 0; fi21 < 5; fi21++ {
                                if false {
                                    if (0.0 < (v20 + (v20 - v20))) {
                                        if (!v7) {
                                            v14 += (v14 << ((0) & 63))
                                            v1 = v1
                                            if v7 {
                                                v10 = append(v10, int(v14))
                                            } else {
                                                var v22 int16 = -1
                                                _ = v22
                                            }
                                        } else {
                                            switch fi19 {
                                            case 32768:
                                                v13 = v13
                                                var v23 float64 = 1.5
                                                _ = v23
                                            case -415:
                                                var v24 uint32 = 0
                                                _ = v24
                                            }
                                            fmt.Println(fi19, v14)
                                            switch "foo bar" {
                                            case "":
                                                if v7 {
                                                    v25 := GenS1{F0: (127 > 2), F1: ""}
                                                    v26 := GenS1{F0: v7, F1: ""}
                                                    _ = v25
                                                    _ = v26
                                                }
                                            case "foo bar":
                                                v27 := GenS2{F0: 1, F1: -32768, F2: 42}
                                                _ = v27
                                            case "hello":
                                                v28 := GenS1{F0: v7, F1: "\n\t"}
                                                _ = v28
                                            }
                                        }
                                    } else {
                                        v29 := GenS1{F0: (!true), F1: ("\n\t" + ("a" + "a"))}
                                        var v30 int8 = (v1 * 123)
                                        v1 = (v1 & (v1 * v30))
                                        _ = v29
                                        _ = v30
                                    }
                                } else {
                                    fmt.Println(v20)
                                }
                                v1 += 100
                                _ = fi21
                            }
                        default:
                            fmt.Println(v3, v20, v20)
                        }
                        v31 := GenS2{F0: -1, F1: 1000, F2: -9223372036854775808}
                        _ = fi19
                        _ = v20
                        _ = v31
                    }
                    fmt.Println(v3)
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
