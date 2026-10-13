package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    if false {
        _ = len(v1)
        if (42 >= 1) {
            var v4 uint8 = 239
            switch ("\n\t" + ((("" + ("a" + "a")) + ((("hello" + "\n\t") + ("\n\t" + "\n\t")) + ("foo bar" + ""))) + ("foo bar" + (("\n\t" + "") + ("a" + ("" + "a")))))) {
            case "a":
                var v5 float64 = float64((float32(1.5) - 3.14159))
                _ = v5
            case "\n\t":
                var v6 bool = (!((2147483647 == 2147483647) || (0.1 != 1e10)))
                _ = v6
            default:
                v7 := GenS1{F0: -1}
                _ = v7
            }
            _ = v4
        } else {
            if true {
                v8 := []int{2147483647, 32768}
                _ = v8
            } else {
                v1[2] = 128
                switch ("foo bar" + (("hello" + ("foo bar" + "")) + ("\n\t" + ("\n\t" + "")))) {
                case "hello":
                    var v9 int32 = -1
                    _ = v9
                case "\n\t":
                    v10 := GenS1{F0: -209}
                    _ = v10
                default:
                    if true {
                        var v11 string = ("hello" + (("" + ((("\n\t" + ("\n\t" + ("a" + "\n\t"))) + "hello") + "\n\t")) + (("" + "\n\t") + "a")))
                        _ = v11
                    }
                }
                v12 := GenS1{F0: -1}
                _ = v12
            }
        }
        var v13 int = -2147483648
        _ = v13
    } else {
        var v14 int8 = 127
        _ = v14
    }
    if (!(2147483647 < 1)) {
        var v15 uint32 = 2147483648
        v15 = v15
        if ((!true) && (32767 >= 1)) {
            switch 0 {
            case -738:
                fmt.Println(v15)
            default:
                var v16 float32 = 1e10
                _ = v16
            }
        } else {
            fmt.Println(v15)
        }
        _ = v15
    }
    if (((((3.14159 / 1e10) - (1e10 - (1e10 + ((0.1 / 1e10) - -2.25)))) <= 0.0) && false) && false) {
        v17 := []int{0, 4294967295, 128}
        _ = v17
    }
    switch 42 {
    case 4294967295:
        v18 := []int{-717, 128, 2147483647}
        _ = v18
    default:
        switch (("foo bar" + "hello") + ("hello" + ("hello" + "foo bar"))) {
        case "":
            v19 := []int{-234, -478, 4294967295}
            _ = v19
        default:
            var v20 int16 = 1000
            _ = v20
        }
    }
    var v21 uint16 = 338
    var v22 int8 = -1
    v21 = v21
    var v23 uint = 42
    var v24 int32 = 2147483647
    fmt.Println(v22, v23)
    var v25 int32 = 2147483647
    var v26 string = "\n\t"
    var v27 bool = false
    for fi28 := 0; fi28 < 12; fi28++ {
        for fi29 := 0; fi29 < 3; fi29++ {
            switch ((v26 + "a") + v26) {
            case "\n\t":
                fmt.Println(fi28, v24, v22)
            case "":
                fmt.Println(v22)
                var v30 int64 = -9223372036854775808
                v31 := []int{(fi29 / ((int(v23)) | 1)), (fi29 << ((v23) & 63))}
                _ = v30
                _ = v31
            case "foo bar":
                v32 := GenS1{F0: 42}
                _ = v32
            }
            if (!false) {
                var v33 uint8 = 127
                _ = v33
            }
            v23 -= 0
            _ = fi29
        }
        v27 = (false && (-1 > 42))
        fmt.Println(v22)
        _ = fi28
    }
    v26 = v26
    switch int(v24) {
    case 972:
        for fi34 := 0; fi34 < 8; fi34++ {
            for fi35 := 0; fi35 < 12; fi35++ {
                v25 = (^(v25))
                v36 := GenS2{F0: v26, F1: 663}
                _ = fi35
                _ = v36
            }
            switch v26 {
            case "foo bar":
                for fi37 := 0; fi37 < 13; fi37++ {
                    var v38 float64 = 1e10
                    switch "a" {
                    case "":
                        fmt.Println(fi34, v27)
                    }
                    if true {
                        v39 := []int{}
                        v22 += 1
                        var v40 int32 = (v24 / ((v24) | 1))
                        _ = v39
                        _ = v40
                    }
                    _ = fi37
                    _ = v38
                }
            case "":
                if false {
                    fmt.Println(v22, v22)
                    v41 := []int{fi34, fi34}
                    _ = v41
                }
            case "hello":
                v42 := []int{127, (fi34 + fi34)}
                _ = v42
            }
            v24 = v25
            _ = fi34
        }
    default:
        v23 = v23
    }
    var v43 int = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v43
}
