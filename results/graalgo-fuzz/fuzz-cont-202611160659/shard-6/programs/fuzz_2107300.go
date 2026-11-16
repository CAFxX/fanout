package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() string {
    switch ("hello" + ("hello" + "foo bar")) {
    case "foo bar":
        var v6 uint64 = 1
        _ = v6
    case "":
        if false {
            v7 := GenS1{F0: 2}
            var v8 int8 = -1
            switch -839 {
            case -1:
                for fi9 := 0; fi9 < 6; fi9++ {
                    var v10 uint8 = 83
                    v10 = v10
                    var v11 uint32 = 65535
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            default:
                v8 = (v8 * v8)
            }
            _ = v7
            _ = v8
        } else {
            var v12 int = 32768
            _ = v12
        }
    }
    var v13 float32 = 1e10
    _ = v13
    return "hello"
}
func genFn1() int16 {
    var v14 uint = 4294967295
    fmt.Println(v14)
    if ((65535 >= 2147483647) || (32767 == 42)) {
        var v15 uint = (v14 % ((2) | 1))
        _ = v15
    }
    _ = v14
    return 0
}

func main() {
    var v1 string = "foo bar"
    var v2 bool = (!((18446744073709551615 <= 729) && true))
    v3 := make(map[int]int)
    var v4 uint = 2
    v5 := make(map[int]uint32)
    if (-38 >= 124) {
        var v16 int32 = -897
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 int32 = (v16 ^ 507)
            if (42 > 0) {
                for fi19 := 0; fi19 < 1; fi19++ {
                    v20 := GenS1{F0: 490}
                    fmt.Println(v1)
                    _ = fi19
                    _ = v20
                }
            }
            _ = fi17
            _ = v18
        }
        switch 128 {
        case 0:
            var v21 uint8 = 2
            _ = v21
        case 4294967295:
            var v22 int32 = (-(v16))
            _ = v22
        }
        _ = v16
    } else {
        if false {
            var v23 uint64 = 0
            v1 = v1
            _ = v23
        } else {
            var v24 int = 973
            v25 := []int{(v24 & 128)}
            _ = v24
            _ = v25
        }
        var v26 int64 = 0
        _ = v26
    }
    var v27 int8 = 123
    v28 := []int{-432, -100}
    v27 -= 0
    v27 = v27
    _ = len(v5)
    var v29 int8 = v27
    v30 := GenS1{F0: 815}
    switch (v1 + v1) {
    case "hello":
        for fi31 := 0; fi31 < 1; fi31++ {
            v29 += v29
            _ = fi31
        }
    }
    if false {
        for fi32 := 0; fi32 < 4; fi32++ {
            if (1 <= -1) {
                var v33 uint16 = 32767
                _ = v33
            } else {
                for fi34 := 0; fi34 < 8; fi34++ {
                    var v35 uint64 = 244
                    _ = fi34
                    _ = v35
                }
                if len(v28) > 0 {
                    v28[0] = -2147483648
                }
            }
            if false {
                v36 := GenS1{F0: 247}
                _ = v36
            } else {
                v37 := GenS1{F0: 0}
                _ = v37
            }
            _ = fi32
        }
        fmt.Println(v29, v2, v2)
        var v38 int16 = -32768
        _ = v38
    }
    v39 := GenS1{F0: 65535}
    var v40 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
