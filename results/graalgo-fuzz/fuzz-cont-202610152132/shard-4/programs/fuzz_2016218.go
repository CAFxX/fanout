package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 uint8 }
func genFn0() uint32 {
    var v7 bool = true
    var v8 float32 = 0.0
    if (v7 || false) {
        v9 := []int{2147483647}
        v10 := GenS1{F0: 1}
        if (((("foo bar" + "a") + "") > "foo bar") || ((v7 && v7) || true)) {
            v11 := GenS1{F0: -544}
            if (!("foo bar" == ("hello" + "a"))) {
                v8 = v8
                v12 := GenS1{F0: 42}
                if len(v9) > 0 {
                    v9[0] = -486
                }
                _ = v12
            } else {
                if len(v9) > 0 {
                    v9[0] = -2147483648
                }
                var v13 int = -146
                _ = v13
            }
            if v7 {
                v7 = (-433 == 9223372036854775807)
                v8 -= v8
            }
            _ = v11
        }
        _ = v10
        _ = v9
    }
    switch "" {
    case "hello":
        var v14 int8 = 42
        _ = v14
    case "a":
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            v15 := GenS2{F0: ""}
            _ = v15
        }
    }
    v7 = v7
    _ = v7
    _ = v8
    return 65535
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    var v4 float32 = 0.1
    v5 := make(map[int]bool)
    var v6 uint32 = 42
    if (-930 >= -575) {
        v16 := []int{-9223372036854775808, 9223372036854775807}
        _ = v16
    }
    if (255 < 255) {
        var v17 int64 = 9223372036854775807
        for fi18 := 0; fi18 < 7; fi18++ {
            v4 = (v4 - float32(0.1))
            _ = fi18
        }
        _ = v17
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        fmt.Println(v1, fi19, v4)
        v20 := []int{}
        _ = fi19
        _ = v20
    }
    fmt.Println(v1)
    switch 42 {
    case -539:
        if false {
            if (true || (952 <= 0)) {
                if (("" + "") >= "a") {
                    var v21 string = ("hello" + (("foo bar" + ("foo bar" + "a")) + "hello"))
                    if (float32((0.1 - 0.1)) > v4) {
                        v21 = (v21 + ((v21 + ((((v21 + v21) + "") + v21) + v21)) + ("a" + v21)))
                        v4 -= 1.5
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 uint32 = 4294967295
                            v24 := []int{fi22, fi22}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                    fmt.Println(v4)
                    _ = v21
                }
                fmt.Println(v6)
                v2[2147483647] = "a"
            }
            v1 = v1
            v25 := GenS1{F0: 1}
            _ = v25
        } else {
            var v26 int32 = -1
            _ = v26
        }
    default:
        var v27 int = 4294967295
        _ = v27
    }
    switch -9223372036854775808 {
    case 32767:
        for fi28 := 0; fi28 < 1; fi28++ {
            fmt.Println(v6)
            v6 = genFn0()
            _ = fi28
        }
    case -2147483648:
        if true {
            var v29 uint64 = uint64(v6)
            _ = v29
        }
    }
    var v30 uint8 = 155
    v31 := GenS2{F0: "hello"}
    var v32 int32 = 42
    var v33 bool = (-108 == -72)
    v34 := GenS3{F0: -97}
    var v35 bool = (v33 && (9223372036854775807 < -77))
    switch ("a" + "\n\t") {
    case "\n\t":
        v33 = ("hello" < "foo bar")
    case "foo bar":
        v30 = v30
    case "":
        v36 := GenS4{F0: 114, F1: (v30 & 42)}
        _ = v36
    default:
        var v37 int8 = -1
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
