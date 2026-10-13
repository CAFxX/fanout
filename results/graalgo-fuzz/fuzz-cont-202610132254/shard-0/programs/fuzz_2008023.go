package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 bool }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 string; F2 string }
type GenS6 struct { F0 string; F1 uint8 }
func genFn0(p0 bool) bool {
    p0 = p0
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{fi7, (fi7 + fi7), fi7}
        _ = fi7
        _ = v8
    }
    var v9 string = "\n\t"
    v10 := GenS1{F0: 255, F1: 775, F2: false}
    _ = p0
    _ = v10
    _ = v9
    return p0
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]bool)
    var v3 uint8 = 90
    var v4 int = 32767
    v5 := make(map[int]int)
    v6 := make(map[uint32]string)
    var v11 int32 = 0
    _ = len(v2)
    var v12 bool = true
    switch 0 {
    case -9223372036854775808:
        switch v4 {
        case -1:
            v13 := GenS1{F0: 1, F1: 65535, F2: (1 == 32767)}
            _ = v13
        default:
            v14 := GenS1{F0: 65535, F1: 2, F2: v12}
            _ = v14
        }
    case 32767:
        switch v4 {
        case 1:
            v3 = (v3 | v3)
        }
    }
    v6[97] = "foo bar"
    v11 = v11
    v15 := GenS2{F0: 2, F1: v12}
    v16 := GenS3{F0: 1, F1: "\n\t"}
    var v17 uint16 = 262
    for fi18 := 0; fi18 < 8; fi18++ {
        if v12 {
            v17 = (v17 - v17)
            switch "\n\t" {
            case "a":
                fmt.Println(v11)
            default:
                var v19 uint64 = 571
                _ = v19
            }
        } else {
            v20 := GenS2{F0: 26, F1: v12}
            _ = v20
        }
        switch ("" + "a") {
        case "a":
            var v21 uint8 = v3
            _ = v21
        case "\n\t":
            if v12 {
                v22 := GenS4{F0: 4294967295}
                if true {
                    for fi23 := 0; fi23 < 10; fi23++ {
                        v3 += (v3 - (v3 - 49))
                        _ = fi23
                    }
                }
                var v24 uint16 = (^(v17))
                _ = v22
                _ = v24
            } else {
                if true {
                    for fi25 := 0; fi25 < 3; fi25++ {
                        delete(v2, uint32(v3))
                        _ = fi25
                    }
                    v26 := GenS5{F0: 112, F1: "a", F2: "\n\t"}
                    if v12 {
                        v27 := GenS1{F0: (^(v17)), F1: 159, F2: (true && v12)}
                        for fi28 := 0; fi28 < 19; fi28++ {
                            v29 := GenS6{F0: "foo bar", F1: v3}
                            _ = fi28
                            _ = v29
                        }
                        var v30 int32 = (v11 << ((2) & 63))
                        _ = v27
                        _ = v30
                    } else {
                        v31 := GenS5{F0: 42, F1: "a", F2: "foo bar"}
                        v4 -= v4
                        v4 = (-(fi18))
                        _ = v31
                    }
                    _ = v26
                }
            }
        }
        _ = fi18
    }
    v32 := v5[-9223372036854775808]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
