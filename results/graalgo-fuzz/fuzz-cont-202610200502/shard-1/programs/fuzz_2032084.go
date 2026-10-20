package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int32; F2 int32 }
type GenS3 struct { F0 uint; F1 uint; F2 uint8 }
func genFn0(p0 uint16, p1 bool, p2 int64) int32 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int16 = 1000
        fmt.Println(p1, p0, p2)
        if (p1 && p1) {
            var v7 bool = false
            _ = v7
        } else {
            v8 := GenS1{F0: -1}
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    switch -1 {
    case 337:
        switch -591 {
        case 2147483647:
            var v9 int32 = 2147483647
            _ = v9
        case 42:
            fmt.Println(p1, p1)
        case 680:
            if true {
                var v10 bool = (false || p1)
                _ = v10
            }
        }
    case 32768:
        if p1 {
            fmt.Println(p1, p2, p0)
        } else {
            if true {
                if (-9223372036854775808 <= 42) {
                    p2 += (p2 & (p2 * 9223372036854775807))
                }
                v11 := GenS1{F0: 100}
                v12 := []int{-2147483648, 32768}
                _ = v11
                _ = v12
            }
        }
    }
    if (p1 && p1) {
        switch ("hello" + ((("foo bar" + "foo bar") + (("\n\t" + "hello") + ((("" + "hello") + "hello") + (("" + "foo bar") + "hello")))) + "foo bar")) {
        case "\n\t":
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := []int{fi13, -517, (fi13 % ((32767) | 1))}
                v15 := GenS2{F0: (^(p0)), F1: 876, F2: 769}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            if (p0 >= (p0 / ((2) | 1))) {
                v17 := GenS3{F0: 392, F1: 42, F2: 78}
                _ = v17
            } else {
                switch (("hello" + "") + ("" + "\n\t")) {
                case "":
                    if (((1000 < 131) || p1) || (p0 >= 2)) {
                        v18 := GenS1{F0: int8(p2)}
                        fmt.Println(p2)
                        _ = v18
                    }
                default:
                    v19 := []int{(fi16 << ((0) & 63)), (fi16 - fi16), fi16, -456}
                    _ = v19
                }
            }
            _ = fi16
        }
    } else {
        fmt.Println(p0, p2)
        var v20 float64 = (1.5 * 1.5)
        p2 -= (p2 / ((0) | 1))
        _ = v20
    }
    var v21 uint = 351
    var v22 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 643
    var v3 uint64 = 1
    var v4 float64 = (1.5 - -2.25)
    fmt.Println(v4)
    if true {
        if false {
            var v23 float32 = 1.5
            _ = v23
        }
    }
    var v24 int = 2147483647
    var v25 int16 = 1
    var v26 uint16 = 32767
    for fi27 := 0; fi27 < 9; fi27++ {
        for fi28 := 0; fi28 < 6; fi28++ {
            if (!true) {
                v24 += (v24 << ((179) & 63))
                _ = len(v1)
            } else {
                v29 := []int{-2147483648}
                fmt.Println(v26, v3)
                _ = v29
            }
            fmt.Println(v26, v24)
            v30 := []int{}
            _ = fi28
            _ = v30
        }
        fmt.Println(v24)
        _ = fi27
    }
    var v31 string = "foo bar"
    v2 += v2
    var v32 uint32 = 4294967295
    for fi33 := 0; fi33 < 4; fi33++ {
        v34 := GenS1{F0: 124}
        _ = fi33
        _ = v34
    }
    var v35 bool = (42 >= 247)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v4
}
