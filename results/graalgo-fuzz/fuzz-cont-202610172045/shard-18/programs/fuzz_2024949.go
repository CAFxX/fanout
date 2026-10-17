package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint8, p2 int16) uint {
    var v7 string = "a"
    v8 := []int{32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int32 = -1
    var v2 bool = false
    var v3 uint16 = 32767
    var v4 string = "\n\t"
    var v5 int32 = v1
    var v6 int16 = 32767
    v9 := []int{127}
    for fi10 := 0; fi10 < 2; fi10++ {
        if v2 {
            if v2 {
                if true {
                    fmt.Println(v6, v2, fi10)
                    fmt.Println(v3, fi10, v2)
                }
                var v11 uint16 = (v3 + 42)
                _ = v11
            }
            if true {
                v9 = append(v9, fi10)
            } else {
                for fi12 := 0; fi12 < 10; fi12++ {
                    fmt.Println(v1)
                    _ = fi12
                }
                for fi13 := 0; fi13 < 20; fi13++ {
                    if v2 {
                        fmt.Println(v4, v5)
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        switch "\n\t" {
                        case "foo bar":
                            switch v4 {
                            case "":
                                var v14 uint16 = v3
                                _ = v14
                            default:
                                var v15 uint16 = 2
                                _ = v15
                            }
                        }
                    }
                    var v16 uint64 = 4294967295
                    var v17 int = 4294967295
                    var v18 uint16 = v3
                    var v19 uint = uint(v18)
                    _ = fi13
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                if len(v9) > 0 {
                    v9[0] = (fi10 | fi10)
                }
            }
        }
        if v2 {
            v5 = (v5 / ((-1) | 1))
            var v20 uint16 = 42
            if (v1 >= (v5 - v1)) {
                if true {
                    var v21 int16 = (v6 | v6)
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v3 = (v20 - (v3 & 0))
                        v5 -= v1
                        v23 := GenS1{F0: -9223372036854775808, F1: 2147483647}
                        v1 += (v1 | v5)
                        v9 = append(v9, fi22)
                        _ = fi22
                        _ = v23
                    }
                    fmt.Println(v4, v3, v4)
                    _ = v21
                } else {
                    v24 := GenS2{F0: (false && v2)}
                    var v25 uint64 = 18446744073709551615
                    _ = v24
                    _ = v25
                }
            }
            _ = v20
        }
        _ = fi10
    }
    v26 := GenS1{F0: 9223372036854775807, F1: 80}
    if v2 {
        _ = len(v9)
    }
    if true {
        if ((180 <= -180) && false) {
            if v2 {
                switch 0 {
                case -867:
                    for fi27 := 0; fi27 < 13; fi27++ {
                        for fi28 := 0; fi28 < 18; fi28++ {
                            v29 := GenS3{F0: (v5 ^ (^(v5)))}
                            v3 = v3
                            _ = fi28
                            _ = v29
                        }
                        var v30 string = ((v4 + (v4 + ((v4 + (v4 + v4)) + (v4 + v4)))) + "foo bar")
                        _ = fi27
                        _ = v30
                    }
                }
            }
            _ = len(v9)
            fmt.Println(v3, v6)
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        var v31 int = 1
        _ = v31
    }
    fmt.Println(v6)
    v6 = 1
    var v32 bool = v2
    v33 := GenS1{F0: 9223372036854775807, F1: 518}
    fmt.Println(v32)
    v34 := GenS4{F0: 0}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v35 uint64 = 0
    v36 := GenS2{F0: v2}
    v1 = 42
    v9 = append(v9, 32767)
    v37 := GenS1{F0: 1, F1: 227}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
