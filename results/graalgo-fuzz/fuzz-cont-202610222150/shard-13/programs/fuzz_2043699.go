package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 50
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: -1}
    switch "a" {
    case "a":
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := v5[88]
            var v9 uint32 = v1
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 uint64 = 88
                var v12 uint8 = (v2 << ((42) & 63))
                v2 = 0
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    case "":
        var v13 float64 = (0.0 + 3.14159)
        _ = v13
    }
    v14 := GenS1{F0: 127}
    fmt.Println(v2, v1)
    v15 := GenS1{F0: 1}
    switch ("hello" + "a") {
    case "\n\t":
        v16 := []int{32768}
        _ = v16
    case "foo bar":
        if true {
            v17 := []int{4294967295, int(v2), 0}
            _ = v17
        } else {
            var v18 float64 = -2.25
            v2 = v2
            v19 := GenS2{F0: true}
            _ = v18
            _ = v19
        }
    default:
        var v20 float64 = 1.5
        _ = v20
    }
    if (!true) {
        v2 += (v2 & v2)
        var v21 int8 = 127
        _ = v21
    } else {
        v2 = v2
        v1 = (v1 & (v1 - (v1 + v1)))
    }
    if true {
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := []int{fi22, (fi22 * fi22)}
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 3; fi24++ {
            if false {
                v1 += v1
                v1 = 65535
            } else {
                v2 = v2
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := []int{-9223372036854775808, (^(fi24)), 127, fi24}
                    switch (("a" + "hello") + (("a" + (("a" + "\n\t") + "hello")) + "\n\t")) {
                    case "hello":
                        v26 = append(v26, fi24)
                    default:
                        v2 += (v2 * 42)
                    }
                    var v27 int = fi25
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 15; fi28++ {
                    v2 += v2
                    _ = fi28
                }
            }
            v29 := GenS2{F0: (false || true)}
            _ = fi24
            _ = v29
        }
    } else {
        var v30 uint = 255
        if (!true) {
            var v31 uint = (v30 * v30)
            _ = v31
        } else {
            if false {
                switch 991 {
                case -22:
                    fmt.Println(v2)
                default:
                    switch "" {
                    case "":
                        v30 = v30
                    }
                }
                for fi32 := 0; fi32 < 11; fi32++ {
                    var v33 uint8 = v2
                    _ = fi32
                    _ = v33
                }
            }
            var v34 int16 = -32768
            _ = v34
        }
        v1 += v1
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
