package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int; F2 uint32 }
func genFn0() uint16 {
    var v4 float32 = 1.5
    v5 := GenS1{F0: 32767, F1: 0}
    v6 := GenS1{F0: 32767, F1: -407}
    var v7 uint64 = 4294967295
    v7 += v7
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 bool, p2 uint64) int32 {
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: 1, F1: (-(p0))}
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 224
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v11 int64 = -9223372036854775808
    switch -422 {
    case -657:
        switch 959 {
        case 4294967295:
            var v12 int = 9223372036854775807
            _ = v12
        case 42:
            for fi13 := 0; fi13 < 11; fi13++ {
                _ = len(v3)
                _ = fi13
            }
        case 142:
            v14 := []int{0, 32768, -2147483648, -826}
            _ = v14
        default:
            fmt.Println(v1, v1)
        }
    case -473:
        v15 := GenS1{F0: 1, F1: 0}
        _ = v15
    case 128:
        var v16 int16 = 356
        _ = v16
    default:
        v11 = v11
    }
    var v17 uint64 = 0
    var v18 string = ""
    if false {
        var v19 float64 = 0.0
        v17 = v17
        _ = v19
    }
    v20 := v3[2147483648]
    v21 := GenS1{F0: 2, F1: -32768}
    switch 32768 {
    case 0:
        if false {
            if true {
                for fi22 := 0; fi22 < 6; fi22++ {
                    v1 = v1
                    v23 := []int{316, fi22, 128}
                    fmt.Println(v18, v17, fi22)
                    _ = fi22
                    _ = v23
                }
            }
            if (42 > 1) {
                switch "foo bar" {
                case "":
                    fmt.Println(v17, v17, v1)
                case "\n\t":
                    fmt.Println(v20, v11)
                }
                v24 := GenS1{F0: 1, F1: -1}
                v1 = (v1 & v1)
                _ = v24
            } else {
                var v25 string = v18
                _ = v25
            }
            switch 127 {
            case -9223372036854775808:
                if (3.14159 != 1e10) {
                    for fi26 := 0; fi26 < 20; fi26++ {
                        var v27 int32 = v1
                        _ = fi26
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 20; fi28++ {
                        if false {
                            v20 = v20
                        } else {
                            v18 = v18
                        }
                        _ = fi28
                    }
                    v29 := GenS2{F0: true}
                    _ = v29
                } else {
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        var v30 uint16 = 2
                        _ = v30
                    case "\n\t":
                        var v31 uint = 65535
                        _ = v31
                    }
                    v17 += v17
                    delete(v3, 1)
                }
            }
        } else {
            var v32 uint = 1
            v33 := GenS2{F0: true}
            v17 = (^(v17))
            _ = v32
            _ = v33
        }
    case 1:
        var v34 int = 24
        _ = v34
    default:
        var v35 uint64 = v17
        _ = v35
    }
    var v36 int16 = -32768
    var v37 int64 = v11
    fmt.Println(v18)
    for fi38 := 0; fi38 < 3; fi38++ {
        for fi39 := 0; fi39 < 14; fi39++ {
            v37 += 0
            switch (v18 + v20) {
            case "hello":
                if true {
                    v40 := v3[0]
                    _ = v40
                }
            }
            v41 := GenS3{F0: 2, F1: (-(fi39)), F2: 938}
            _ = fi39
            _ = v41
        }
        _ = fi38
    }
    var v42 uint64 = v17
    switch (v18 + ("foo bar" + "foo bar")) {
    case "a":
        var v43 int64 = v37
        _ = v43
    case "\n\t":
        v44 := GenS2{F0: ((true && false) && (380 >= 42))}
        _ = v44
    case "hello":
        var v45 int64 = v11
        _ = v45
    }
    v46 := []int{721, int(v1)}
    if false {
        var v47 int64 = 998
        _ = v47
    }
    var v48 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v42
    _ = v46
    _ = v48
}
