package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int64 }
type GenS2 struct { F0 int }
func genFn0() uint32 {
    var v7 int16 = 0
    fmt.Println(v7)
    _ = v7
    return 715
}

func main() {
    var v1 uint16 = 915
    var v2 uint64 = 701
    v3 := []string{"foo bar", "a", "", "\n\t"}
    var v4 string = "a"
    var v5 uint64 = v2
    v6 := []int{-2147483648, 42, 948, -1}
    v1 = (v1 % (((v1 >> ((0) & 63))) | 1))
    var v8 int16 = -32768
    var v9 int32 = 831
    v1 = v1
    switch -9223372036854775808 {
    case 128:
        var v10 uint = 0
        _ = v10
    case 2147483647:
        var v11 int = 9223372036854775807
        _ = v11
    }
    var v12 string = ""
    fmt.Println(v4, v1)
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint8 = 138
        var v15 uint32 = 4294967295
        v16 := GenS1{F0: 0, F1: -2147483648, F2: 1}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    switch v4 {
    case "foo bar":
        v17 := GenS1{F0: 127, F1: 0, F2: 1}
        _ = v17
    case "hello":
        v2 -= 866
    case "\n\t":
        fmt.Println(v12)
    default:
        for fi18 := 0; fi18 < 4; fi18++ {
            if true {
                var v19 bool = false
                var v20 uint32 = 791
                var v21 uint = 784
                _ = v19
                _ = v20
                _ = v21
            } else {
                v4 = ((v4 + "a") + "a")
                fmt.Println(v9, v4, v4)
            }
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: 42, F1: (v9 >> ((759) & 63)), F2: -422}
                v6 = append(v6, 187)
                v8 -= v8
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v4 = v12
                _ = fi22
                _ = v23
            }
            _ = fi18
        }
    }
    switch 32767 {
    case -9223372036854775808:
        if ((v9 + (v9 + (v9 * v9))) < 0) {
            for fi24 := 0; fi24 < 9; fi24++ {
                v1 -= v1
                for fi25 := 0; fi25 < 18; fi25++ {
                    switch fi25 {
                    case 1:
                        _ = len(v3)
                    }
                    if true {
                        v26 := GenS2{F0: fi25}
                        var v27 int8 = -67
                        if len(v3) > 0 {
                            v3[0] = v4
                        }
                        _ = v26
                        _ = v27
                    }
                    _ = fi25
                }
                _ = fi24
            }
        }
    case 148:
        if (11 >= 2147483647) {
            fmt.Println(v12, v8)
        } else {
            v12 = ""
            fmt.Println(v1, v12, v12)
        }
    case 32767:
        v2 -= (v2 * 645)
    }
    if false {
        v6 = append(v6, 9223372036854775807)
        var v28 uint = 65535
        _ = v28
    }
    _ = len(v3)
    v8 = (v8 >> ((1) & 63))
    v6 = append(v6, 32767)
    v12 = (v12 + v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
