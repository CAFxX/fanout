package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    v2 := []string{("" + "foo bar")}
    v3 := make(map[uint32]int)
    var v4 uint16 = 0
    var v5 bool = false
    fmt.Println(v4)
    var v6 int32 = 1
    switch "\n\t" {
    case "hello":
        switch 127 {
        case 1:
            if v5 {
                for fi7 := 0; fi7 < 18; fi7++ {
                    delete(v1, (fi7 << ((262) & 63)))
                    _ = fi7
                }
                for fi8 := 0; fi8 < 3; fi8++ {
                    var v9 int8 = 127
                    if false {
                        var v10 int = (fi8 ^ 42)
                        for fi11 := 0; fi11 < 9; fi11++ {
                            v6 -= (v6 - v6)
                            _ = fi11
                        }
                        _ = v10
                    }
                    var v12 int32 = int32(v9)
                    _ = fi8
                    _ = v12
                    _ = v9
                }
                var v13 int32 = (v6 + (v6 * -119))
                _ = v13
            } else {
                v6 -= (v6 >> ((18446744073709551615) & 63))
                if len(v2) > 0 {
                    v2[0] = (("\n\t" + (("" + "hello") + (("\n\t" + "") + "\n\t"))) + "hello")
                }
                fmt.Println(v4, v6, v4)
            }
        }
    case "foo bar":
        if v5 {
            var v14 int16 = -935
            v15 := v1[-1]
            _ = v14
            _ = v15
        } else {
            var v16 int = 0
            var v17 uint = 990
            var v18 int8 = 1
            _ = v16
            _ = v17
            _ = v18
        }
    case "":
        if v5 {
            var v19 float32 = 1e10
            v5 = v5
            _ = v19
        }
    }
    v5 = v5
    v20 := GenS1{F0: 1000}
    v21 := GenS1{F0: -1}
    v6 -= v6
    v4 -= (v4 - (v4 | (v4 + v4)))
    for fi22 := 0; fi22 < 1; fi22++ {
        if v5 {
            v23 := GenS2{F0: (336 != 8)}
            var v24 float64 = -2.25
            _ = v23
            _ = v24
        }
        if ((fi22 % (((fi22 >> ((872) & 63))) | 1)) == 9223372036854775807) {
            fmt.Println(v5, fi22)
            for fi25 := 0; fi25 < 9; fi25++ {
                switch "a" {
                case "\n\t":
                    v6 += v6
                default:
                    v6 = (v6 << ((0) & 63))
                }
                if v5 {
                    if (!(-32768 == 0)) {
                        var v26 string = "foo bar"
                        v5 = v5
                        _ = v26
                    }
                }
                v3[4294967295] = fi25
                fmt.Println(fi25)
                if v5 {
                    v4 -= v4
                    v27 := GenS3{F0: 9223372036854775807}
                    _ = v27
                } else {
                    _ = len(v2)
                    if v5 {
                        var v28 uint = 150
                        _ = v28
                    }
                }
                _ = fi25
            }
            var v29 int32 = (v6 % ((0) | 1))
            _ = v29
        }
        if true {
            v4 = 32767
            var v30 int = fi22
            _ = v30
        }
        _ = fi22
    }
    var v31 int32 = v6
    var v32 float32 = (3.14159 * -2.25)
    v31 = v6
    if (!v5) {
        v32 += (v32 + 1e10)
    }
    fmt.Println(v5)
    v33 := GenS2{F0: (42 < 0)}
    v6 -= 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
