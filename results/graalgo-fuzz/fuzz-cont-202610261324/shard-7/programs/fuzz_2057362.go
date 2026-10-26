package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int8 }

func main() {
    v1 := []int{}
    var v2 string = "foo bar"
    v3 := []uint32{855, 647, 42, 0}
    v4 := make(map[string]string)
    switch "foo bar" {
    case "foo bar":
        v2 = "a"
    case "\n\t":
        for fi5 := 0; fi5 < 6; fi5++ {
            v6 := GenS1{F0: (fi5 - fi5)}
            _ = fi5
            _ = v6
        }
    case "hello":
        var v7 float32 = ((3.14159 / 0.1) + (3.14159 * 1e10))
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 bool = (!true)
            v10 := GenS1{F0: (fi8 >> ((353) & 63))}
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(v2)
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int32 = 0
            for fi13 := 0; fi13 < 15; fi13++ {
                v14 := v4[v2]
                switch (fi11 + fi13) {
                case 9223372036854775807:
                    fmt.Println(v14, v12)
                }
                if false {
                    var v15 int8 = 87
                    _ = v15
                } else {
                    v1 = append(v1, fi13)
                    var v16 int64 = 754
                    var v17 uint8 = 80
                    _ = v16
                    _ = v17
                }
                _ = fi13
                _ = v14
            }
            _ = fi11
            _ = v12
        }
    }
    v2 = (v2 + ("" + "foo bar"))
    v18 := GenS2{F0: 653}
    v19 := GenS3{F0: 236, F1: 100}
    var v20 float32 = (((0.0 - ((1e10 / (-2.25 / 1e10)) / -2.25)) / float32(1.5)) + float32(1.5))
    v21 := v4[(v2 + v2)]
    for fi22 := 0; fi22 < 12; fi22++ {
        v3 = append(v3, 4294967295)
        var v23 uint32 = 911
        _ = fi22
        _ = v23
    }
    fmt.Println(v21, v2)
    switch (("\n\t" + ("" + (v2 + "\n\t"))) + (v2 + ((v2 + v2) + ""))) {
    case "foo bar":
        var v24 int = -9223372036854775808
        _ = v24
    case "":
        if (false || (!false)) {
            if false {
                v2 = "a"
                v21 = v21
            }
            for fi25 := 0; fi25 < 1; fi25++ {
                v2 = v21
                switch -9223372036854775808 {
                case -263:
                    v26 := GenS1{F0: (fi25 | fi25)}
                    _ = v26
                case -9223372036854775808:
                    v20 += (((0.0 / (1e10 - -2.25)) - v20) * (v20 * (v20 * (-2.25 / (v20 + float32(1e10))))))
                default:
                    for fi27 := 0; fi27 < 10; fi27++ {
                        switch fi25 {
                        case 535:
                            switch (v2 + (("\n\t" + v21) + "\n\t")) {
                            case "\n\t":
                                if true {
                                    if ((((((v2 + ("" + v2)) + v21) + "foo bar") + "foo bar") + v21) <= ("a" + "\n\t")) {
                                        if (42 == 38) {
                                            var v28 uint = 0
                                            _ = v28
                                        }
                                        fmt.Println(v20, fi25, v2)
                                    } else {
                                        v21 = ("a" + "hello")
                                        v29 := GenS2{F0: 42}
                                        _ = v29
                                    }
                                    if (((-2.25 * 0.0) > (3.14159 - (3.14159 + 3.14159))) && false) {
                                        v30 := GenS2{F0: 1}
                                        v20 = 0.1
                                        if false {
                                            delete(v4, "")
                                            v31 := GenS3{F0: 255, F1: -1}
                                            _ = v31
                                        } else {
                                            if (0 < 255) {
                                                v32 := GenS2{F0: 934}
                                                _ = v32
                                            }
                                        }
                                        _ = v30
                                    }
                                }
                            case "a":
                                v2 = v2
                            }
                        case -308:
                            v33 := GenS1{F0: (fi27 >> ((939) & 63))}
                            _ = v33
                        default:
                            v2 = "\n\t"
                        }
                        var v34 int32 = 721
                        _ = fi27
                        _ = v34
                    }
                }
                _ = fi25
            }
        }
    default:
        var v35 string = ((v21 + "\n\t") + v2)
        _ = v35
    }
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
