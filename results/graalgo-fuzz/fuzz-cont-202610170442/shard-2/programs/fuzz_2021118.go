package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 int16; F2 bool }
func genFn0() int16 {
    var v7 int = 4294967295
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int64 = 127
        if true {
            v9 = v9
            var v10 int = 193
            var v11 bool = (-2147483648 >= 42)
            _ = v10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v12 := []int{(^(v7)), v7}
    v13 := GenS1{F0: 0}
    for fi14 := 0; fi14 < 2; fi14++ {
        v7 -= (fi14 | (v7 * -869))
        fmt.Println(v7, fi14)
        _ = fi14
    }
    _ = v12
    _ = v13
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    var v3 uint16 = 174
    var v4 uint = uint(v3)
    var v5 bool = (119 > 42)
    var v6 int16 = 0
    delete(v1, "hello")
    v6 += 1
    v6 = -989
    v15 := GenS1{F0: v6}
    delete(v1, "hello")
    if v5 {
        if (true && v5) {
            if (true && true) {
                fmt.Println(v6, v6, v6)
                var v16 int16 = genFn0()
                _ = v16
            }
            for fi17 := 0; fi17 < 12; fi17++ {
                v18 := GenS2{F0: -1}
                fmt.Println(fi17)
                _ = fi17
                _ = v18
            }
        }
        if v5 {
            v19 := GenS3{F0: -1}
            var v20 uint8 = 1
            if true {
                v6 = (v6 | v6)
            } else {
                var v21 uint8 = (v20 - (v20 >> ((v4) & 63)))
                v4 += v4
                var v22 int32 = 0
                _ = v21
                _ = v22
            }
            _ = v19
            _ = v20
        }
    } else {
        switch 501 {
        case 4294967295:
            fmt.Println(v6, v3, v5)
        case 32768:
            v4 = 396
        case -963:
            v4 += v4
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint8 = 1
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v3, v6)
            _ = fi23
            _ = v24
        }
    }
    if (1 <= -9223372036854775808) {
        delete(v1, "hello")
        v2 = append(v2, -865)
    } else {
        var v25 uint16 = (v3 * uint16(v4))
        var v26 float64 = 1.5
        _ = v25
        _ = v26
    }
    if (true && v5) {
        if false {
            for fi27 := 0; fi27 < 12; fi27++ {
                v1[(("hello" + "\n\t") + "")] = int(v4)
                _ = fi27
            }
        }
        v6 = v6
        v4 = 4294967295
    } else {
        var v28 uint16 = (v3 & 42)
        _ = v28
    }
    v5 = (0.0 >= 3.14159)
    switch ("foo bar" + "foo bar") {
    case "a":
        v29 := GenS4{F0: 42}
        _ = v29
    }
    var v30 int16 = (v6 / ((-32768) | 1))
    for fi31 := 0; fi31 < 18; fi31++ {
        var v32 bool = v5
        v32 = v32
        _ = fi31
        _ = v32
    }
    v5 = (-9223372036854775808 >= 1)
    if (1 <= (v30 >> (((v4 | (v4 * 755))) & 63))) {
        for fi33 := 0; fi33 < 10; fi33++ {
            v3 = v3
            _ = fi33
        }
        v34 := GenS2{F0: 42}
        switch "a" {
        case "a":
            v35 := GenS3{F0: 127}
            _ = v35
        case "hello":
            v36 := v1["hello"]
            _ = v36
        }
        _ = v34
    }
    if len(v2) > 0 {
        v2[0] = 127
    }
    if true {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            v4 -= (^(v4))
        case "hello":
            if false {
                var v37 float64 = (-2.25 / (((1e10 + float64(((0.1 + ((1e10 + 1.5) + -2.25)) - 1e10))) * 0.1) / 1.5))
                var v38 uint16 = v3
                var v39 int16 = v30
                _ = v37
                _ = v38
                _ = v39
            }
        case "":
            _ = len(v2)
        }
        v40 := GenS5{F0: 1, F1: v30, F2: v5}
        if v5 {
            v4 = (v4 + 529)
        } else {
            if v5 {
                fmt.Println(v3, v5)
                var v41 string = ("" + ("\n\t" + "foo bar"))
                _ = v41
            }
            v4 -= uint(v3)
        }
        _ = v40
    } else {
        var v42 int64 = -295
        fmt.Println(v3, v4)
        _ = v42
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        v3 -= v3
        v6 = genFn0()
        _ = len(v2)
        _ = fi43
    }
    var v44 uint = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
