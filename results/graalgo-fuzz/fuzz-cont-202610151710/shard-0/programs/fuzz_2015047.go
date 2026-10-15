package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int32; F2 int32 }
type GenS3 struct { F0 uint8; F1 string; F2 int64 }

func main() {
    var v1 int64 = 1
    v2 := []int{}
    var v3 int64 = -1
    v4 := make(map[string]int)
    var v5 uint16 = 188
    for fi6 := 0; fi6 < 10; fi6++ {
        switch fi6 {
        case 32768:
            v2 = append(v2, fi6)
        case -179:
            var v7 uint = 2
            _ = v7
        case 1:
            for fi8 := 0; fi8 < 18; fi8++ {
                v5 += 2
                v9 := GenS1{F0: 38}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                if true {
                    delete(v4, ("a" + ""))
                    var v10 uint64 = 1
                    _ = v10
                } else {
                    fmt.Println(fi6)
                    delete(v4, "")
                }
                if (!false) {
                    _ = len(v2)
                    _ = len(v2)
                } else {
                    if (false && true) {
                        var v11 float64 = (3.14159 + 1.5)
                        _ = v11
                    }
                    var v12 int16 = -678
                    _ = v12
                }
                if false {
                    switch fi8 {
                    case 4294967295:
                        fmt.Println(v1, fi8)
                    case 0:
                        v5 = v5
                    case 32768:
                        _ = len(v2)
                    }
                } else {
                    v5 = v5
                    fmt.Println(v1, fi6, v3)
                }
                _ = fi8
                _ = v9
            }
        default:
            _ = len(v2)
        }
        if false {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 float32 = 0.0
                var v15 int = fi6
                _ = fi13
                _ = v14
                _ = v15
            }
            v1 = 1
        }
        var v16 uint64 = 2
        _ = fi6
        _ = v16
    }
    fmt.Println(v1)
    fmt.Println(v1)
    if true {
        var v17 string = "\n\t"
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 12; fi19++ {
                _ = len(v4)
                fmt.Println(v5)
                if len(v2) > 0 {
                    v2[0] = (fi18 - fi19)
                }
                _ = fi19
            }
            _ = fi18
        }
        _ = v17
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v21 := GenS2{F0: 42, F1: 1, F2: 453}
        _ = fi20
        _ = v21
    }
    switch "\n\t" {
    case "hello":
        if (127 > 0) {
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 int = fi22
                if ((!(v3 < (v3 << ((255) & 63)))) || (uint32(v5) >= 125)) {
                    fmt.Println(v23)
                    var v24 uint8 = 0
                    fmt.Println(v23)
                    _ = v24
                } else {
                    if true {
                        var v25 float64 = 0.1
                        for fi26 := 0; fi26 < 11; fi26++ {
                            var v27 uint8 = 1
                            fmt.Println(v23, v5)
                            fmt.Println(v5)
                            _ = fi26
                            _ = v27
                        }
                        fmt.Println(v23)
                        _ = v25
                    }
                }
                for fi28 := 0; fi28 < 11; fi28++ {
                    if (false || (4294967295 < 127)) {
                        v23 = (fi22 & fi28)
                    }
                    _ = fi28
                }
                _ = fi22
                _ = v23
            }
            for fi29 := 0; fi29 < 18; fi29++ {
                v30 := GenS3{F0: 42, F1: "", F2: 1}
                var v31 float64 = 0.1
                var v32 int16 = 42
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
        }
    default:
        v33 := GenS1{F0: 127}
        _ = v33
    }
    v34 := GenS2{F0: 1000, F1: -471, F2: 2147483647}
    var v35 uint16 = v5
    v36 := GenS3{F0: 4, F1: "a", F2: 126}
    var v37 int16 = 32767
    if false {
        for fi38 := 0; fi38 < 12; fi38++ {
            v35 = v5
            var v39 uint64 = 4294967295
            _ = fi38
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
