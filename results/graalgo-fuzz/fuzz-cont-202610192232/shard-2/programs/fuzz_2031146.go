package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int8; F1 uint; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0(p0 string) int {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{(fi5 & -1), fi5}
        if len(v6) > 0 {
            v6[0] = fi5
        }
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int8 = 42
        _ = fi7
        _ = v8
    }
    if true {
        var v9 float32 = 1e10
        if true {
            var v10 uint32 = 813
            _ = v10
        }
        _ = v9
    }
    _ = p0
    return -1000
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{127}
    v3 := []string{"a", ((("\n\t" + "foo bar") + (("hello" + ("hello" + "a")) + "foo bar")) + ("hello" + "foo bar")), "hello", ""}
    var v4 int64 = -9223372036854775808
    v4 = -1
    v11 := GenS1{F0: 128, F1: 0}
    _ = len(v1)
    for fi12 := 0; fi12 < 16; fi12++ {
        if false {
            fmt.Println(v4, fi12)
            v13 := v1[65535]
            _ = v13
        }
        var v14 uint64 = 18446744073709551615
        var v15 int8 = -115
        _ = fi12
        _ = v14
        _ = v15
    }
    var v16 uint8 = 2
    if ((uint32(v16) != 2147483648) && false) {
        for fi17 := 0; fi17 < 6; fi17++ {
            delete(v1, 412)
            _ = len(v1)
            _ = fi17
        }
        var v18 int64 = -142
        _ = v18
    } else {
        v16 = (v16 ^ (v16 & 0))
    }
    v4 = v4
    v1[4294967295] = "hello"
    v19 := GenS1{F0: (v16 - v16), F1: -1}
    if true {
        if false {
            for fi20 := 0; fi20 < 12; fi20++ {
                v4 += v4
                _ = fi20
            }
            var v21 float64 = 0.1
            _ = v21
        }
    } else {
        var v22 uint32 = 2
        fmt.Println(v16, v22)
        v16 += v16
        _ = v22
    }
    v23 := GenS2{F0: 78, F1: 1, F2: 75}
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 uint64 = 367
        if true {
            v4 = (^(v4))
            if true {
                switch (fi24 ^ 128) {
                case 127:
                    var v26 uint64 = (^(v25))
                    _ = v26
                case 42:
                    if true {
                        fmt.Println(v25)
                        v2 = append(v2, fi24)
                    } else {
                        switch ("foo bar" + ("a" + "\n\t")) {
                        case "":
                            fmt.Println(fi24, v16)
                        }
                    }
                }
            } else {
                var v27 int16 = 32767
                fmt.Println(fi24, v25)
                fmt.Println(v27, v16)
                _ = v27
            }
            if (fi24 == -191) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v16 += v16
                v16 += v16
            }
        } else {
            fmt.Println(v4, v4, v4)
        }
        if false {
            var v28 uint64 = 18446744073709551615
            if false {
                switch "\n\t" {
                case "a":
                    var v29 int64 = v4
                    _ = v29
                case "foo bar":
                    v25 = 2
                case "hello":
                    fmt.Println(v16)
                }
                var v30 uint64 = (v25 & (v28 ^ (v28 - v28)))
                switch ("" + "foo bar") {
                case "foo bar":
                    switch fi24 {
                    case 1:
                        var v31 int16 = -32768
                        _ = v31
                    default:
                        v32 := GenS3{F0: 4294967295}
                        _ = v32
                    }
                }
                _ = v30
            } else {
                for fi33 := 0; fi33 < 18; fi33++ {
                    v28 -= 0
                    v3 = append(v3, "")
                    _ = fi33
                }
                var v34 bool = ((v16 >= v16) && (true && false))
                _ = v34
            }
            _ = v28
        } else {
            v35 := GenS4{F0: (-(fi24)), F1: 0}
            for fi36 := 0; fi36 < 6; fi36++ {
                if ((!true) || true) {
                    _ = len(v3)
                }
                var v37 int64 = v4
                var v38 float64 = (((0.0 - 0.0) - 1.5) - float64(3.14159))
                _ = fi36
                _ = v37
                _ = v38
            }
            _ = v35
        }
        _ = fi24
        _ = v25
    }
    if (660 < 2) {
        v39 := v1[197]
        _ = v39
    }
    switch "" {
    case "\n\t":
        switch "a" {
        case "hello":
            v40 := GenS1{F0: 10, F1: 0}
            _ = v40
        case "\n\t":
            v4 -= 9223372036854775807
        case "":
            var v41 bool = (-2.25 == -2.25)
            _ = v41
        }
    }
    v42 := GenS4{F0: genFn0(("a" + ("\n\t" + ("hello" + "foo bar")))), F1: 42}
    v16 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v42
}
