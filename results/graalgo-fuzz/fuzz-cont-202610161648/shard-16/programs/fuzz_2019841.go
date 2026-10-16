package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 int32, p1 bool, p2 uint8) int32 {
    p0 = p0
    var v6 int16 = 0
    var v7 int8 = 64
    v8 := GenS1{F0: 2, F1: 18446744073709551615, F2: 255}
    switch -461 {
    case 42:
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS1{F0: 1, F1: 2, F2: 647}
            fmt.Println(fi9)
            p0 = (p0 << ((42) & 63))
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (p0 & (p0 << ((525) & 63)))
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 2147483648
    var v3 int32 = 2147483647
    v4 := make(map[uint32]int)
    v5 := []int{-9223372036854775808, 867}
    var v11 int = -1
    if (false || (-509 > 0)) {
        if true {
            v12 := GenS2{F0: true}
            _ = v12
        } else {
            fmt.Println(v11, v3, v3)
            if (42 <= 0) {
                if false {
                    _ = len(v4)
                    fmt.Println(v11)
                    v2 = 42
                } else {
                    v11 -= (v11 / ((32767) | 1))
                }
                if false {
                    var v13 uint = 255
                    if true {
                        var v14 uint32 = 1
                        _ = v14
                    } else {
                        v15 := v4[(v2 + (v2 | 2))]
                        _ = v15
                    }
                    v16 := GenS2{F0: true}
                    _ = v13
                    _ = v16
                }
            }
        }
    }
    var v17 float32 = (3.14159 + (3.14159 * 1e10))
    for fi18 := 0; fi18 < 17; fi18++ {
        if (0 == 387) {
            fmt.Println(v11, fi18, v3)
            v19 := GenS1{F0: (v2 * v2), F1: 352, F2: 0}
            var v20 int = v11
            _ = v19
            _ = v20
        } else {
            v17 -= v17
            switch fi18 {
            case 9223372036854775807:
                v2 = (v2 | 65535)
            }
        }
        v21 := v4[v2]
        _ = fi18
        _ = v21
    }
    v5 = append(v5, -661)
    var v22 uint = 335
    v23 := GenS1{F0: (v2 / ((v2) | 1)), F1: 643, F2: 1}
    v22 -= 18446744073709551615
    switch (v11 & 811) {
    case 32767:
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS2{F0: false}
            _ = fi24
            _ = v25
        }
    }
    if len(v5) > 0 {
        v5[0] = (v11 + (v11 * v11))
    }
    switch ((("\n\t" + "") + ("\n\t" + ("hello" + ""))) + "") {
    case "":
        var v26 string = ("foo bar" + ("foo bar" + "hello"))
        _ = v26
    default:
        switch 9223372036854775807 {
        case 299:
            if (!(-1 <= 32767)) {
                if true {
                    var v27 int32 = v3
                    v27 = v3
                    _ = v27
                } else {
                    delete(v1, 2147483648)
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v22 -= (v22 + v22)
            } else {
                v28 := GenS3{F0: (!false), F1: (v3 / ((v3) | 1))}
                v29 := GenS1{F0: (^(v2)), F1: 0, F2: 2}
                _ = v28
                _ = v29
            }
        }
    }
    v30 := GenS2{F0: (!false)}
    var v31 int32 = (^(v3))
    if (4294967295 <= (v2 << ((v22) & 63))) {
        v2 += 65535
        switch "foo bar" {
        case "hello":
            switch 42 {
            case -1:
                var v32 uint32 = (^(v2))
                _ = v32
            }
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                var v33 bool = (42 <= -9223372036854775808)
                _ = v33
            }
        }
    }
    var v34 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
