package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 string; F2 uint64 }
type GenS4 struct { F0 uint16 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 int = fi4
        fmt.Println(fi4, v5)
        v6 := []int{fi4, 32767}
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 4294967295}
    _ = v7
    return 760
}
func genFn1() uint {
    if false {
        var v8 uint32 = 2147483648
        var v9 float64 = 3.14159
        v10 := []int{-350, 1, 128, 1}
        _ = v10
        _ = v8
        _ = v9
    } else {
        if (genFn0() == 1) {
            switch -434 {
            case 42:
                switch 1 {
                case -794:
                    v11 := []int{-389}
                    _ = v11
                }
            }
        } else {
            if (3.14159 <= 1.5) {
                var v12 int16 = 416
                switch (("" + ("foo bar" + "hello")) + "a") {
                case "hello":
                    v13 := GenS2{F0: 4294967295}
                    _ = v13
                case "\n\t":
                    v14 := GenS1{F0: 594}
                    _ = v14
                case "":
                    v12 += 757
                }
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS3{F0: 2, F1: ("foo bar" + ("hello" + "a")), F2: 2}
                    v17 := GenS4{F0: 42}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = v12
            } else {
                v18 := []int{631, 4294967295, 989}
                var v19 int16 = 1000
                _ = v18
                _ = v19
            }
        }
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 int64 = -1
            _ = fi20
            _ = v21
        }
        v22 := []int{}
        _ = v22
    }
    var v23 int16 = -1
    v24 := []int{444, 127, 127}
    for fi25 := 0; fi25 < 17; fi25++ {
        if true {
            var v26 uint = 1
            var v27 uint32 = 2
            _ = v26
            _ = v27
        } else {
            if (true && true) {
                var v28 int16 = v23
                var v29 bool = (((!false) && ((("a" + "hello") + "a") == "\n\t")) && (-9223372036854775808 < 1))
                fmt.Println(fi25, fi25, v29)
                _ = v28
                _ = v29
            } else {
                v24 = append(v24, fi25)
                var v30 uint64 = 0
                _ = v30
            }
            v23 += (^(v23))
            var v31 float32 = 3.14159
            _ = v31
        }
        v23 = (-(v23))
        fmt.Println(v23)
        _ = fi25
    }
    _ = v23
    _ = v24
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 687
    var v3 uint8 = 111
    var v32 uint16 = 65535
    fmt.Println(v3, v32, v3)
    v1[9223372036854775807] = ("hello" + ("foo bar" + ("" + "foo bar")))
    v3 += (v3 / (((^(v3))) | 1))
    v33 := GenS3{F0: 42, F1: (("\n\t" + "\n\t") + ("\n\t" + "hello")), F2: 0}
    v3 -= v3
    for fi34 := 0; fi34 < 6; fi34++ {
        switch (fi34 | (fi34 | fi34)) {
        case -1:
            v35 := GenS4{F0: v32}
            _ = v35
        case 301:
            v36 := GenS2{F0: 42}
            _ = v36
        }
        _ = fi34
    }
    var v37 string = "foo bar"
    var v38 float64 = ((0.1 * 1.5) - 0.1)
    if false {
        for fi39 := 0; fi39 < 5; fi39++ {
            v2 = -993
            v38 -= v38
            var v40 uint = 255
            _ = fi39
            _ = v40
        }
        v1[32768] = v37
        fmt.Println(v3, v3, v37)
    }
    for fi41 := 0; fi41 < 14; fi41++ {
        for fi42 := 0; fi42 < 12; fi42++ {
            _ = len(v1)
            v2 = -1
            _ = fi42
        }
        var v43 string = "\n\t"
        v44 := GenS1{F0: 18446744073709551615}
        _ = fi41
        _ = v43
        _ = v44
    }
    for fi45 := 0; fi45 < 6; fi45++ {
        v37 = v37
        var v46 int32 = 304
        _ = fi45
        _ = v46
    }
    v47 := GenS1{F0: 81}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v38
    _ = v47
}
