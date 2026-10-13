package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -2147483648
    v3 := make(map[string]uint32)
    var v4 string = (("a" + "\n\t") + "hello")
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 18; fi6++ {
            var v7 uint16 = 1
            _ = fi6
            _ = v7
        }
        var v8 uint16 = 634
        v4 = (v4 + "foo bar")
        _ = fi5
        _ = v8
    }
    var v9 uint8 = 42
    switch 1 {
    case 42:
        delete(v3, v4)
    case 374:
        var v10 int64 = -1
        _ = v10
    }
    var v11 uint8 = v9
    var v12 uint = 2
    var v13 int64 = -9223372036854775808
    v14 := []int{4294967295}
    v2 = v2
    if false {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 float32 = (3.14159 / -2.25)
            _ = fi15
            _ = v16
        }
        var v17 uint64 = 280
        if (!true) {
            v12 = v12
        }
        _ = v17
    } else {
        if true {
            v4 = ((((((((v4 + (v4 + (v4 + (((("foo bar" + "hello") + (("\n\t" + "") + "hello")) + "a") + "a")))) + "a") + "\n\t") + "\n\t") + (("a" + "foo bar") + v4)) + "foo bar") + "a") + v4)
            fmt.Println(v12, v13)
            var v18 int = 501
            _ = v18
        } else {
            fmt.Println(v9)
        }
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        _ = len(v1)
        var v20 string = v4
        _ = fi19
        _ = v20
    }
    if true {
        v21 := v1[v4]
        _ = v21
    } else {
        var v22 uint16 = 285
        var v23 uint64 = 4294967295
        v3[v4] = 65535
        _ = v22
        _ = v23
    }
    if ((32767 > 124) || false) {
        if false {
            if (!false) {
                if len(v14) > 0 {
                    v14[0] = -9223372036854775808
                }
                var v24 int8 = 116
                _ = v24
            }
            for fi25 := 0; fi25 < 12; fi25++ {
                v26 := GenS1{F0: (("hello" + ("" + ("\n\t" + ("hello" + v4)))) + v4), F1: (-(v2))}
                _ = fi25
                _ = v26
            }
            v27 := GenS1{F0: "foo bar", F1: (v2 / ((-2147483648) | 1))}
            _ = v27
        } else {
            v28 := GenS1{F0: v4, F1: -1}
            _ = v28
        }
        fmt.Println(v2, v11)
    } else {
        v12 = (v12 >> ((406) & 63))
        var v29 int32 = v2
        v13 -= v13
        _ = v29
    }
    if true {
        var v30 float64 = (3.14159 / 1.0)
        _ = v30
    }
    if len(v14) > 0 {
        v14[0] = 128
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        v13 = (v13 | -1)
        if (!false) {
            v32 := v1[""]
            fmt.Println(v4, v4, v13)
            switch ("hello" + "\n\t") {
            case "foo bar":
                fmt.Println(v11, fi31)
            case "a":
                var v33 int32 = 715
                _ = v33
            }
            _ = v32
        }
        _ = fi31
    }
    switch 42 {
    case 32767:
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 int32 = v2
            var v36 uint64 = 2
            v37 := GenS1{F0: (("hello" + (v4 + "hello")) + v4), F1: (v2 & (v35 / ((v2) | 1)))}
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
        }
    case 9223372036854775807:
        var v38 bool = ((false && (v11 >= (^(v11)))) && true)
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
