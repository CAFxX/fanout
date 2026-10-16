package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int8; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 uint16 = 392
    v4 = 1
    v4 += v4
    if ((!true) || ((v4 % (((v4 << ((65535) & 63))) | 1)) == 230)) {
        for fi5 := 0; fi5 < 4; fi5++ {
            v6 := []int{}
            _ = fi5
            _ = v6
        }
    }
    if (!true) {
        var v7 string = "foo bar"
        v4 = (v4 % ((32767) | 1))
        v8 := GenS1{F0: -2147483648}
        _ = v7
        _ = v8
    }
    v4 -= (^(v4))
    v4 = 708
    var v9 int64 = -318
    for fi10 := 0; fi10 < 20; fi10++ {
        switch "foo bar" {
        case "foo bar":
            if ("hello" >= "\n\t") {
                v11 := GenS1{F0: -118}
                v9 += v9
                if (42 <= 127) {
                    if ((v9 >> ((65535) & 63)) == (-(v9))) {
                        var v12 int8 = -66
                        switch (fi10 & -9223372036854775808) {
                        case 528:
                            for fi13 := 0; fi13 < 2; fi13++ {
                                v14 := []int{}
                                _ = fi13
                                _ = v14
                            }
                        }
                        v12 = v12
                        _ = v12
                    }
                    v4 += (v4 >> ((0) & 63))
                    v9 += (v9 * v9)
                }
                _ = v11
            }
        case "\n\t":
            var v15 uint32 = 1
            _ = v15
        }
        _ = fi10
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        if (0 != 2147483648) {
            if (1000 <= -320) {
                v17 := []int{(^(fi16)), -1, (-(fi16)), -108}
                v9 = (v9 - v9)
                _ = v17
            }
            v18 := []int{fi16, (fi16 - (-(fi16))), (fi16 << ((2) & 63)), fi16}
            if (0 >= -729) {
                if false {
                    v4 = 1
                } else {
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(fi16)
                    }
                    var v19 int8 = 1
                    v9 -= v9
                    _ = v19
                }
                var v20 uint64 = 42
                _ = v20
            }
            _ = v18
        } else {
            v4 += (v4 << ((1) & 63))
            v9 = v9
            var v21 uint = 2
            _ = v21
        }
        _ = fi16
    }
    fmt.Println(v9)
    var v22 uint64 = 1
    v23 := GenS1{F0: -2147483648}
    v9 -= (v9 >> ((835) & 63))
    _ = len(v3)
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v9, fi25, fi24)
            v1[2147483648] = (2 < 128)
            _ = fi25
        }
        if true {
            var v26 int = (^(fi24))
            v26 += fi24
            _ = v26
        } else {
            v27 := GenS1{F0: 2147483647}
            var v28 uint = 845
            var v29 uint16 = (^(v4))
            _ = v27
            _ = v28
            _ = v29
        }
        _ = fi24
    }
    var v30 string = "a"
    var v31 int64 = (v9 >> ((389) & 63))
    v32 := GenS2{F0: v4, F1: -103, F2: v31}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v9
}
