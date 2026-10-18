package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 int = -9223372036854775808
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 float64 = float64((3.14159 + float32(1e10)))
        var v7 int = 2147483647
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 float64 = ((1e10 + 1e10) * -2.25)
    var v9 string = "\n\t"
    for fi10 := 0; fi10 < 5; fi10++ {
        if (!(!true)) {
            var v11 string = v9
            for fi12 := 0; fi12 < 18; fi12++ {
                v9 = v11
                v13 := v2[-1]
                v14 := []int{(-(v3)), fi12}
                var v15 int8 = -128
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            v8 = 0.0
            _ = v11
        }
        _ = fi10
    }
    fmt.Println(v9)
    var v16 uint32 = 888
    var v17 uint64 = 700
    _ = len(v4)
    v8 = 3.14159
    if false {
        if false {
            switch v9 {
            case "":
                v16 = (v16 >> ((740) & 63))
            case "a":
                if (false && (!false)) {
                    fmt.Println(v8, v9)
                    var v18 float32 = 1.5
                    var v19 int16 = 42
                    _ = v18
                    _ = v19
                }
            default:
                if (!true) {
                    v20 := GenS1{F0: 100}
                    var v21 uint8 = 0
                    for fi22 := 0; fi22 < 19; fi22++ {
                        switch ("a" + "a") {
                        case "a":
                            v23 := []int{}
                            _ = v23
                        case "hello":
                            var v24 uint32 = (v16 + (v16 - 885))
                            _ = v24
                        case "":
                            fmt.Println(v8, fi22)
                        }
                        _ = fi22
                    }
                    _ = v20
                    _ = v21
                }
            }
        } else {
            v9 = ""
        }
    } else {
        if (false && false) {
            var v25 int64 = 42
            _ = v25
        } else {
            var v26 uint64 = v17
            _ = v26
        }
        fmt.Println(v3, v16)
        if (false || false) {
            v8 = v8
            for fi27 := 0; fi27 < 11; fi27++ {
                for fi28 := 0; fi28 < 6; fi28++ {
                    v29 := GenS2{F0: v17, F1: 2}
                    if (-468 < 1) {
                        var v30 int32 = 42
                        v31 := []int{(fi28 ^ fi27), (v3 / (((fi27 - (^(v3)))) | 1)), (v3 >> ((18446744073709551615) & 63))}
                        var v32 int8 = -128
                        _ = v30
                        _ = v31
                        _ = v32
                    } else {
                        var v33 int = -938
                        var v34 uint64 = 238
                        _ = v33
                        _ = v34
                    }
                    v17 += 182
                    _ = fi28
                    _ = v29
                }
                if false {
                    var v35 uint16 = 2
                    var v36 int32 = 88
                    _ = v35
                    _ = v36
                } else {
                    if true {
                        switch (v9 + (v9 + (v9 + "\n\t"))) {
                        case "foo bar":
                            _ = len(v2)
                        }
                        var v37 float32 = -2.25
                        _ = v37
                    }
                }
                switch 0 {
                case 1:
                    switch v9 {
                    case "hello":
                        v38 := []int{(fi27 % ((9223372036854775807) | 1))}
                        _ = v38
                    case "":
                        var v39 uint8 = 2
                        _ = v39
                    case "foo bar":
                        v9 = v9
                    }
                default:
                    fmt.Println(v16, v9, fi27)
                }
                _ = fi27
            }
            var v40 float32 = -2.25
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        if true {
            v8 = v8
            if (11 <= 1) {
                fmt.Println(v3, v16, fi41)
                v42 := []int{(fi41 / ((2147483647) | 1))}
                v3 += fi41
                _ = v42
            } else {
                fmt.Println(v3, v8, v3)
                v9 = "\n\t"
            }
        }
        _ = fi41
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        fmt.Println(v16)
        if (!false) {
            for fi44 := 0; fi44 < 7; fi44++ {
                if true {
                    v3 = fi44
                    v8 = -2.25
                    var v45 uint64 = (v17 % (((^(v17))) | 1))
                    _ = v45
                }
                var v46 int16 = 0
                _ = fi44
                _ = v46
            }
        }
        if (v17 == v17) {
            var v47 string = v9
            _ = v47
        }
        _ = fi43
    }
    if true {
        var v48 int = (v3 >> ((0) & 63))
        v49 := GenS3{F0: 0, F1: (v17 & 4294967295), F2: 0}
        _ = v48
        _ = v49
    }
    var v50 int = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v8
    _ = v9
}
