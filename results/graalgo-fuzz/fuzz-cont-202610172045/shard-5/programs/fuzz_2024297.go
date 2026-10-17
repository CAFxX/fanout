package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 int32 }
func genFn0() int64 {
    switch -2147483648 {
    case -74:
        var v5 int8 = -76
        _ = v5
    case -9223372036854775808:
        var v6 int = 4294967295
        _ = v6
    default:
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint = 0
            var v9 uint32 = 2
            switch fi7 {
            case 330:
                var v10 int16 = 1000
                _ = v10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v11 int = 1
    var v12 bool = (30 < -52)
    var v13 bool = true
    if (549 != 2147483647) {
        var v14 int16 = -946
        for fi15 := 0; fi15 < 2; fi15++ {
            if v13 {
                v13 = (v12 && ((!true) || (1 < 2)))
                switch "" {
                case "\n\t":
                    var v16 bool = false
                    _ = v16
                case "foo bar":
                    v14 += v14
                }
                fmt.Println(v14, v11)
            }
            _ = fi15
        }
        _ = v14
    } else {
        var v17 float64 = ((((3.14159 - 0.1) + 0.0) - ((1.5 + (0.1 - float64(((0.1 + 0.1) / float32(-2.25))))) - float64((((0.1 / 1.0) + -2.25) / 3.14159)))) / 1.0)
        v11 = -9223372036854775808
        _ = v17
    }
    _ = v11
    _ = v12
    _ = v13
    return -703
}
func genFn1() string {
    v18 := []int{1, -255}
    switch 9223372036854775807 {
    case -1:
        var v19 string = "foo bar"
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        if false {
            fmt.Println(fi20)
            v21 := GenS1{F0: 689, F1: -2147483648, F2: 255}
            _ = v21
        }
        _ = fi20
    }
    var v22 int64 = -609
    _ = v18
    _ = v22
    return (("hello" + "hello") + "\n\t")
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]int)
    var v3 string = "foo bar"
    var v4 int8 = -1
    v23 := GenS1{F0: 42, F1: 0, F2: 127}
    fmt.Println(v3)
    if (1 < 1) {
        fmt.Println(v3, v1)
        v24 := GenS2{F0: false, F1: 117, F2: -757}
        _ = v24
    } else {
        if false {
            v4 = (v4 ^ v4)
            switch 128 {
            case 42:
                var v25 int64 = (v1 - 0)
                _ = v25
            case 32768:
                var v26 bool = false
                _ = v26
            case -516:
                v4 = v4
            default:
                v27 := []int{402, 32768}
                _ = v27
            }
            for fi28 := 0; fi28 < 18; fi28++ {
                v29 := []int{454, fi28, (-(fi28))}
                var v30 int64 = (-(v1))
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        var v31 uint16 = 669
        for fi32 := 0; fi32 < 9; fi32++ {
            v33 := GenS1{F0: 751, F1: 903, F2: 255}
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    if false {
        if (100 < 127) {
            v34 := GenS2{F0: (1e10 >= (1e10 + (1.5 + -2.25))), F1: 2147483648, F2: -1}
            fmt.Println(v3)
            var v35 int8 = (v4 >> ((592) & 63))
            _ = v34
            _ = v35
        }
        var v36 float32 = 3.14159
        _ = v36
    }
    var v37 bool = true
    v38 := []int{-9223372036854775808}
    v39 := GenS1{F0: 2, F1: -1, F2: 2}
    v1 = (v1 ^ v1)
    for fi40 := 0; fi40 < 20; fi40++ {
        for fi41 := 0; fi41 < 6; fi41++ {
            var v42 uint8 = 255
            var v43 int32 = 1
            _ = fi41
            _ = v42
            _ = v43
        }
        _ = fi40
    }
    var v44 string = ""
    switch ("foo bar" + v44) {
    case "":
        var v45 string = ("hello" + ("foo bar" + v3))
        _ = v45
    case "hello":
        v44 = ""
    }
    v46 := GenS1{F0: 0, F1: 0, F2: 0}
    switch 151 {
    case 9223372036854775807:
        var v47 float32 = -2.25
        _ = v47
    case 42:
        switch (v44 + v44) {
        case "foo bar":
            var v48 uint64 = 18446744073709551615
            _ = v48
        }
    case 2147483647:
        fmt.Println(v37, v1, v44)
    }
    v37 = true
    switch 2147483647 {
    case 910:
        v49 := GenS1{F0: 0, F1: 0, F2: 42}
        _ = v49
    }
    if len(v38) > 0 {
        v38[0] = -220
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v44
    _ = v46
}
