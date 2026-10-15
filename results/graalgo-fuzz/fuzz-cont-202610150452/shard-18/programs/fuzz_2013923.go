package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int32 }
func genFn0() int16 {
    var v4 uint16 = 2
    fmt.Println(v4)
    switch 127 {
    case -696:
        fmt.Println(v4)
    }
    switch "\n\t" {
    case "foo bar":
        v4 = v4
    }
    _ = v4
    return 614
}

func main() {
    var v1 string = ""
    var v2 uint = 2
    v3 := make(map[int]uint32)
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: v2, F1: 1}
        v7 := v3[fi5]
        switch fi5 {
        case 194:
            var v8 bool = false
            _ = v8
        case -189:
            _ = len(v3)
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        switch 32767 {
        case -2147483648:
            v1 = v1
        default:
            for fi9 := 0; fi9 < 12; fi9++ {
                switch (("" + "hello") + "foo bar") {
                case "hello":
                    v1 = ""
                case "a":
                    switch v1 {
                    case "\n\t":
                        v3[32767] = 2147483648
                    case "a":
                        var v10 int16 = 1
                        _ = v10
                    }
                default:
                    fmt.Println(fi9)
                }
                var v11 int8 = -120
                switch (fi9 << ((2) & 63)) {
                case 4294967295:
                    switch fi9 {
                    case -695:
                        v1 = (((v1 + ("" + v1)) + v1) + v1)
                    case 32768:
                        if (!false) {
                            switch 127 {
                            case 2147483647:
                                if (1 <= 0) {
                                    v1 = (v1 + "\n\t")
                                    v3[128] = 4294967295
                                }
                            }
                            var v12 uint64 = 0
                            _ = v12
                        } else {
                            v13 := []int{fi9}
                            v14 := v3[fi9]
                            _ = v13
                            _ = v14
                        }
                    case 9223372036854775807:
                        var v15 bool = (0 > 1)
                        _ = v15
                    }
                default:
                    v16 := v3[fi9]
                    _ = v16
                }
                _ = fi9
                _ = v11
            }
        }
    } else {
        switch "foo bar" {
        case "":
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 uint64 = 418
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        }
    }
    var v20 uint16 = 255
    v21 := GenS1{F0: (v2 / ((v2) | 1)), F1: -128}
    var v22 uint16 = 32767
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint32 = 790
        fmt.Println(v2, v24, v2)
        _ = fi23
        _ = v24
    }
    if (false || false) {
        v2 = v2
    }
    if true {
        delete(v3, 1)
        var v25 uint16 = 2
        v26 := GenS1{F0: v2, F1: 0}
        _ = v25
        _ = v26
    } else {
        var v27 uint = v2
        var v28 uint64 = 42
        v29 := GenS2{F0: 0, F1: v22, F2: uint32(v28)}
        _ = v27
        _ = v28
        _ = v29
    }
    v30 := []int{}
    _ = len(v3)
    var v31 uint32 = 586
    v22 = 0
    if len(v30) > 0 {
        v30[0] = 0
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 uint32 = (v31 - v31)
        v34 := GenS1{F0: v2, F1: -128}
        var v35 int32 = 668
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        switch v1 {
        case "":
            v2 = (v2 - v2)
        default:
            v37 := GenS3{F0: 0}
            _ = v37
        }
        _ = fi36
    }
    if ((^(v2)) > v2) {
        var v38 uint64 = 0
        var v39 int = int(v38)
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
}
