package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int; F2 uint32 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int8 = -59
    v2 := make(map[int]bool)
    var v3 int8 = v1
    var v4 uint = 789
    v5 := make(map[string]bool)
    var v6 int16 = 439
    switch ("hello" + ("hello" + "\n\t")) {
    case "hello":
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint = (^(v4))
            _ = fi7
            _ = v8
        }
    case "":
        fmt.Println(v1, v6)
    default:
        v6 = (v6 | 871)
    }
    switch ("foo bar" + (("\n\t" + "") + (("" + "a") + "hello"))) {
    case "\n\t":
        var v9 int32 = -1
        _ = v9
    case "hello":
        if true {
            if (false || true) {
                switch (("foo bar" + "\n\t") + "a") {
                case "a":
                    v10 := []int{676, -9223372036854775808, 42}
                    _ = v10
                default:
                    v11 := GenS1{F0: v4}
                    _ = v11
                }
                var v12 float32 = (-2.25 * 1.5)
                v6 -= v6
                _ = v12
            } else {
                switch "hello" {
                case "hello":
                    fmt.Println(v3, v4, v6)
                }
                var v13 bool = true
                v14 := GenS1{F0: (v4 / (((v4 - v4)) | 1))}
                _ = v13
                _ = v14
            }
            fmt.Println(v3)
            v15 := GenS1{F0: 850}
            _ = v15
        }
    case "a":
        var v16 int = 9223372036854775807
        _ = v16
    }
    v3 += (v1 | v1)
    switch int(v4) {
    case -2147483648:
        v4 = v4
    default:
        fmt.Println(v3, v6)
    }
    var v17 string = "\n\t"
    var v18 int64 = 42
    v17 = v17
    var v19 uint64 = 206
    var v20 uint64 = v19
    v19 = 4294967295
    fmt.Println(v6, v20, v6)
    if (v19 < 42) {
        v3 += (v3 >> ((v4) & 63))
    } else {
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 uint16 = 303
            switch (fi21 / (((fi21 % ((-869) | 1))) | 1)) {
            case 137:
                for fi23 := 0; fi23 < 19; fi23++ {
                    fmt.Println(v3)
                    var v24 uint32 = 65535
                    _ = fi23
                    _ = v24
                }
            case -658:
                v4 += 2
            default:
                v25 := GenS1{F0: 18446744073709551615}
                _ = v25
            }
            _ = fi21
            _ = v22
        }
        if (v3 < 42) {
            fmt.Println(v4, v18)
        } else {
            fmt.Println(v17, v17, v18)
        }
        v3 -= (v3 | (v1 + (v3 | -45)))
    }
    if ((int8(v18) != v3) || (337 >= 65535)) {
        var v26 int32 = 1
        for fi27 := 0; fi27 < 10; fi27++ {
            fmt.Println(v18)
            fmt.Println(v20, v26)
            v28 := []int{}
            _ = fi27
            _ = v28
        }
        _ = v26
    } else {
        var v29 uint8 = 255
        delete(v2, -144)
        _ = v29
    }
    v30 := GenS2{F0: -926, F1: 128, F2: 0}
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 int16 = v6
        if false {
            if ((v18 << ((865) & 63)) == v18) {
                v33 := []int{(fi31 % ((-2147483648) | 1)), (fi31 - (^(fi31)))}
                _ = v33
            } else {
                v4 = 4294967295
            }
        }
        var v34 string = v17
        _ = fi31
        _ = v32
        _ = v34
    }
    v35 := GenS3{F0: (^(v4))}
    var v36 uint32 = 1
    v37 := v5[v17]
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
