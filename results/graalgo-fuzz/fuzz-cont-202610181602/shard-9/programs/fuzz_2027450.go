package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int32) bool {
    var v7 uint64 = 1
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(fi8)
            _ = fi9
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            switch "hello" {
            case "a":
                if (449 >= 255) {
                    v11 := GenS1{F0: (("\n\t" + "\n\t") + ""), F1: -32768, F2: 62}
                    _ = v11
                } else {
                    fmt.Println(fi8)
                }
            case "\n\t":
                switch "hello" {
                case "hello":
                    var v12 bool = (374 >= 42)
                    _ = v12
                case "\n\t":
                    var v13 bool = false
                    _ = v13
                default:
                    if (42 >= (p0 - p0)) {
                        switch "" {
                        case "a":
                            var v14 uint = 65535
                            _ = v14
                        case "\n\t":
                            v15 := GenS1{F0: "a", F1: -146, F2: 42}
                            _ = v15
                        case "":
                            p1 = p1
                        default:
                            v16 := GenS1{F0: "a", F1: -1, F2: -1}
                            _ = v16
                        }
                    } else {
                        v17 := GenS2{F0: (!true), F1: fi8}
                        _ = v17
                    }
                }
            }
            _ = fi10
        }
        _ = fi8
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v7
    return true
}
func genFn1(p0 int8) bool {
    switch 42 {
    case 2147483647:
        if (!((false || true) && (2147483647 >= 2147483647))) {
            for fi18 := 0; fi18 < 13; fi18++ {
                p0 -= (p0 / (((p0 + 100)) | 1))
                _ = fi18
            }
        }
    default:
        v19 := GenS3{F0: 255}
        _ = v19
    }
    p0 = (p0 << ((65535) & 63))
    v20 := GenS4{F0: 255}
    p0 = 127
    p0 = p0
    _ = p0
    _ = v20
    return (4294967295 >= 379)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 33
    var v3 int8 = -84
    var v4 uint16 = 195
    v5 := make(map[string]int)
    var v6 uint = 42
    var v21 int16 = -125
    fmt.Println(v3)
    v6 = v2
    v22 := GenS1{F0: ("\n\t" + ("hello" + ("hello" + ""))), F1: v21, F2: 127}
    if false {
        var v23 float32 = 1.5
        _ = v23
    }
    v2 = v6
    v24 := GenS2{F0: true, F1: -973}
    if false {
        v25 := []int{}
        if false {
            v26 := GenS3{F0: (v4 / ((0) | 1))}
            _ = v26
        }
        _ = v25
    } else {
        _ = len(v5)
        var v27 uint8 = 255
        fmt.Println(v21, v4, v2)
        _ = v27
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 uint64 = 1
        v30 := []int{(fi28 + (fi28 - int(v4))), fi28, fi28}
        _ = fi28
        _ = v29
        _ = v30
    }
    v2 -= v2
    v21 += -1
    if true {
        if true {
            _ = len(v1)
            v31 := v5["hello"]
            _ = v31
        } else {
            var v32 int = 461
            _ = v32
        }
        fmt.Println(v4, v4, v21)
        v33 := GenS1{F0: "foo bar", F1: v21, F2: v3}
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
