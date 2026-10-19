package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int; F2 bool }
func genFn0() int64 {
    v6 := GenS1{F0: (0 <= 1)}
    if true {
        switch 75 {
        case 42:
            if false {
                if true {
                    v7 := GenS1{F0: true}
                    v8 := GenS1{F0: ((42 > 18446744073709551615) && true)}
                    _ = v7
                    _ = v8
                }
                var v9 float32 = (1.5 * 1e10)
                _ = v9
            }
        case -2147483648:
            var v10 string = "a"
            _ = v10
        case 128:
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 uint = 10
                _ = fi11
                _ = v12
            }
        }
        var v13 string = (("foo bar" + "foo bar") + "hello")
        _ = v13
    } else {
        var v14 int16 = -1
        var v15 uint16 = 32767
        v16 := []int{235, 4294967295}
        _ = v14
        _ = v15
        _ = v16
    }
    switch "a" {
    case "":
        var v17 float32 = -2.25
        _ = v17
    case "\n\t":
        v18 := []int{969, 32767, 32767}
        _ = v18
    default:
        var v19 int16 = 42
        _ = v19
    }
    var v20 uint64 = 2
    v21 := []int{2147483647, 4294967295, 1, 33}
    _ = v20
    _ = v21
    _ = v6
    return 1
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]bool)
    var v3 bool = true
    v4 := make(map[int]bool)
    var v5 int64 = -1
    var v22 int64 = (v5 & (v5 * 1))
    v22 = (v22 + (v5 + genFn0()))
    var v23 string = "a"
    switch int(v22) {
    case 794:
        v24 := v4[9223372036854775807]
        _ = v24
    default:
        var v25 string = v23
        _ = v25
    }
    var v26 int16 = -1
    var v27 int64 = 0
    fmt.Println(v26, v5, v26)
    var v28 bool = ((("\n\t" + ((v23 + (((v23 + ("a" + v23)) + (v23 + "\n\t")) + "a")) + "")) + ((v23 + "hello") + (v23 + "\n\t"))) != "foo bar")
    v22 -= -1
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := GenS2{F0: (v22 % (((v5 + 42)) | 1)), F1: (^(fi29)), F2: v3}
        _ = fi29
        _ = v30
    }
    v3 = true
    if v3 {
        switch 1 {
        case 9223372036854775807:
            switch v23 {
            case "foo bar":
                switch ("" + v23) {
                case "foo bar":
                    fmt.Println(v23, v23, v22)
                }
            case "\n\t":
                for fi31 := 0; fi31 < 18; fi31++ {
                    v27 -= v5
                    v3 = v3
                    v5 -= v27
                    _ = fi31
                }
            default:
                switch 42 {
                case 9223372036854775807:
                    v32 := GenS2{F0: v27, F1: 997, F2: (v23 < v23)}
                    _ = v32
                case 2147483647:
                    v28 = v28
                case 659:
                    var v33 bool = false
                    _ = v33
                }
            }
        }
        v27 -= v27
        switch ((v23 + "foo bar") + "") {
        case "hello":
            switch "a" {
            case "":
                if (127 == 255) {
                    v26 -= 0
                    _ = len(v4)
                    fmt.Println(v23, v5)
                }
            case "\n\t":
                v28 = (!false)
            }
        case "a":
            fmt.Println(v22)
        case "":
            var v34 uint = 980
            _ = v34
        }
    } else {
        switch -1 {
        case 127:
            _ = len(v4)
        case 2147483647:
            if v3 {
                var v35 int32 = -2147483648
                fmt.Println(v5)
                _ = v35
            } else {
                fmt.Println(v27, v26)
                delete(v2, -9223372036854775808)
                v36 := []int{}
                _ = v36
            }
        }
        var v37 string = (v23 + v23)
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
