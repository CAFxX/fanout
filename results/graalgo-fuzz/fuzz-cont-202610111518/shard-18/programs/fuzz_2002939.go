package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint8 = 75
    v2 := []int{-613, 1}
    var v3 int16 = -992
    var v4 string = ("" + "")
    var v5 uint = 24
    if len(v2) > 0 {
        v2[0] = 1
    }
    fmt.Println(v5)
    if false {
        switch (v4 + v4) {
        case "foo bar":
            var v6 int64 = 741
            _ = v6
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch 4294967295 {
        case 2147483647:
            v2 = append(v2, -844)
        }
    }
    if (1 > 935) {
        switch 2147483647 {
        case -9223372036854775808:
            if true {
                for fi7 := 0; fi7 < 19; fi7++ {
                    v3 = (v3 << ((42) & 63))
                    var v8 uint64 = 18446744073709551615
                    var v9 string = v4
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                for fi10 := 0; fi10 < 10; fi10++ {
                    v3 = (v3 ^ v3)
                    _ = fi10
                }
                if false {
                    if len(v2) > 0 {
                        v2[0] = int(v3)
                    }
                    v1 = v1
                    for fi11 := 0; fi11 < 15; fi11++ {
                        fmt.Println(fi11)
                        var v12 float64 = 1e10
                        v12 -= 0.1
                        _ = fi11
                        _ = v12
                    }
                }
            } else {
                fmt.Println(v5, v4)
            }
        case 127:
            v4 = "\n\t"
        case -309:
            if (!true) {
                fmt.Println(v5, v4, v1)
                fmt.Println(v1, v4)
                fmt.Println(v3, v4)
            } else {
                v2 = append(v2, 32768)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        default:
            fmt.Println(v4, v1)
        }
    } else {
        switch -1 {
        case 9223372036854775807:
            if false {
                v2 = append(v2, 376)
                switch v4 {
                case "a":
                    _ = len(v2)
                case "hello":
                    fmt.Println(v4)
                case "foo bar":
                    v13 := GenS1{F0: 3, F1: 2147483648}
                    _ = v13
                default:
                    v14 := GenS1{F0: 51, F1: 254}
                    _ = v14
                }
            } else {
                v4 = (v4 + v4)
                v15 := GenS2{F0: 32768, F1: 127}
                var v16 uint64 = 18446744073709551615
                _ = v15
                _ = v16
            }
        case -1:
            v17 := GenS2{F0: 32768, F1: -2147483648}
            _ = v17
        default:
            for fi18 := 0; fi18 < 7; fi18++ {
                v3 += 32767
                _ = fi18
            }
        }
        var v19 int64 = int64(v1)
        _ = v19
    }
    fmt.Println(v4, v3)
    var v20 int32 = 0
    for fi21 := 0; fi21 < 20; fi21++ {
        v3 -= v3
        _ = fi21
    }
    _ = len(v2)
    v20 += (^(v20))
    if false {
        var v22 uint16 = 42
        _ = v22
    }
    fmt.Println(v4, v4, v4)
    switch 589 {
    case 0:
        var v23 uint32 = 453
        _ = v23
    case -22:
        switch 42 {
        case 1:
            if len(v2) > 0 {
                v2[0] = 2147483647
            }
        default:
            for fi24 := 0; fi24 < 19; fi24++ {
                for fi25 := 0; fi25 < 2; fi25++ {
                    v4 = v4
                    _ = fi25
                }
                _ = fi24
            }
        }
    default:
        for fi26 := 0; fi26 < 4; fi26++ {
            v27 := GenS1{F0: 42, F1: 42}
            _ = fi26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := GenS1{F0: 100, F1: 65535}
        _ = fi28
        _ = v29
    }
    v30 := GenS3{F0: 567, F1: (v20 << ((v5) & 63))}
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 uint64 = 2
        v33 := GenS3{F0: 0, F1: v20}
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
