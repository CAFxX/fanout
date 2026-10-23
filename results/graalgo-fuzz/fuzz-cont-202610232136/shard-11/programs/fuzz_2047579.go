package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint16) string {
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: fi6, F1: 1}
        v8 := []int{fi6, -1, fi6}
        var v9 uint8 = 70
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        switch (^(fi10)) {
        case -1:
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := []int{}
                _ = fi11
                _ = v12
            }
        default:
            for fi13 := 0; fi13 < 17; fi13++ {
                p0 += 896
                _ = fi13
            }
        }
        var v14 uint = 255
        _ = fi10
        _ = v14
    }
    _ = p0
    return ("" + "hello")
}
func genFn1(p0 string, p1 uint32, p2 int) bool {
    if true {
        if false {
            switch (p2 ^ -1) {
            case 4294967295:
                fmt.Println(p1, p2, p0)
            case -387:
                p0 = (p0 + "a")
            default:
                p1 = p1
            }
            for fi15 := 0; fi15 < 19; fi15++ {
                p0 = (p0 + "a")
                _ = fi15
            }
        }
    } else {
        p0 = ""
        var v16 uint64 = 253
        _ = v16
    }
    var v17 int8 = 77
    v18 := GenS1{F0: (p2 | p2), F1: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return false
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    var v3 int8 = -81
    var v4 uint16 = 1
    v5 := []string{"", "foo bar"}
    var v19 bool = (255 != 42)
    v19 = (v19 || (-801 == 42))
    var v20 string = "hello"
    switch (v20 + v20) {
    case "hello":
        if v19 {
            var v21 bool = v19
            var v22 string = "foo bar"
            v20 = (v22 + v22)
            _ = v21
            _ = v22
        } else {
            switch 0 {
            case 2147483647:
                var v23 int64 = 985
                _ = v23
            }
            if v19 {
                var v24 int64 = 9223372036854775807
                if true {
                    fmt.Println(v19)
                } else {
                    for fi25 := 0; fi25 < 15; fi25++ {
                        fmt.Println(fi25, v19)
                        _ = fi25
                    }
                    v24 += v24
                    v3 = -1
                }
                _ = v24
            }
        }
    case "foo bar":
        if ((int8(v4) != (v3 * 1)) || (1e10 > -2.25)) {
            v20 = v20
            if (v19 && (18446744073709551615 != 1)) {
                if false {
                    switch -564 {
                    case 32767:
                        fmt.Println(v4, v19)
                    case 4294967295:
                        fmt.Println(v19, v3, v20)
                    }
                    v20 = ((v20 + (("\n\t" + "") + "foo bar")) + v20)
                    v20 = v20
                } else {
                    for fi26 := 0; fi26 < 11; fi26++ {
                        _ = len(v5)
                        _ = fi26
                    }
                    v4 += v4
                    if len(v5) > 0 {
                        v5[0] = v20
                    }
                }
                v19 = v19
            }
            fmt.Println(v20)
        }
    default:
        v27 := GenS2{F0: v19, F1: (v4 ^ (v4 << ((29) & 63)))}
        _ = v27
    }
    fmt.Println(v19)
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := v2[10]
        for fi30 := 0; fi30 < 15; fi30++ {
            var v31 uint16 = v4
            _ = fi30
            _ = v31
        }
        var v32 int64 = -1
        _ = fi28
        _ = v29
        _ = v32
    }
    v5 = append(v5, v20)
    switch "a" {
    case "hello":
        v33 := GenS1{F0: 127, F1: -201}
        _ = v33
    case "a":
        v34 := GenS2{F0: v19, F1: 65535}
        _ = v34
    case "foo bar":
        v3 -= v3
    default:
        var v35 uint64 = 399
        _ = v35
    }
    var v36 int16 = 42
    v37 := GenS1{F0: 0, F1: -2147483648}
    var v38 uint = 196
    _ = len(v5)
    var v39 uint64 = 1
    switch -576 {
    case -821:
        var v40 int8 = int8(v39)
        _ = v40
    default:
        for fi41 := 0; fi41 < 14; fi41++ {
            switch v20 {
            case "hello":
                v4 = v4
            case "":
                v42 := GenS1{F0: 37, F1: 42}
                _ = v42
            case "a":
                var v43 uint = uint(v4)
                _ = v43
            }
            _ = fi41
        }
    }
    if false {
        var v44 bool = (2147483647 < 9223372036854775807)
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
