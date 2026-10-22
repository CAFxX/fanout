package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    v4 := make(map[uint32]bool)
    var v5 float32 = 3.14159
    v6 := make(map[uint32]bool)
    v1 += v1
    v7 := []int{42, 1}
    switch ("a" + "") {
    case "a":
        fmt.Println(v1, v1)
    case "hello":
        if true {
            if len(v7) > 0 {
                v7[0] = -798
            }
        } else {
            var v8 int32 = (v1 % ((-810) | 1))
            v9 := GenS1{F0: 655}
            _ = v8
            _ = v9
        }
    case "\n\t":
        v2[-179] = 65535
    default:
        var v10 int = -1
        _ = v10
    }
    v1 += (v1 + (-(v1)))
    fmt.Println(v1)
    fmt.Println(v5, v1)
    var v11 int32 = 203
    if (807 <= -309) {
        if (-9223372036854775808 == 706) {
            if (!((2147483648 < 2) || false)) {
                if true {
                    v12 := v6[951]
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        _ = len(v7)
                    case "\n\t":
                        v13 := v6[1]
                        _ = v13
                    }
                    fmt.Println(v1, v11)
                    _ = v12
                }
            }
        } else {
            v1 += -805
            switch "" {
            case "":
                fmt.Println(v11, v11)
            case "a":
                if false {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        var v15 int16 = -1
                        _ = fi14
                        _ = v15
                    }
                } else {
                    fmt.Println(v1)
                    v16 := v3[-247]
                    _ = v16
                }
            }
        }
        if (!(v11 <= v1)) {
            for fi17 := 0; fi17 < 14; fi17++ {
                v11 += v11
                _ = fi17
            }
        }
    } else {
        v7 = append(v7, 202)
    }
    if ((!(!true)) && false) {
        var v18 uint16 = 42
        v11 = 2147483647
        _ = v18
    } else {
        v1 = v11
        switch 42 {
        case 32768:
            if (!true) {
                fmt.Println(v11)
                if false {
                    var v19 uint64 = 0
                    if true {
                        var v20 uint8 = 42
                        delete(v6, 0)
                        _ = v20
                    } else {
                        if (-1 == 456) {
                            v11 = 0
                        }
                        var v21 int = 42
                        _ = v21
                    }
                    var v22 uint16 = 65535
                    _ = v19
                    _ = v22
                }
                if (!false) {
                    var v23 int32 = v1
                    var v24 string = "foo bar"
                    _ = v23
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 9; fi25++ {
                        v26 := GenS2{F0: fi25, F1: 42}
                        v11 -= v11
                        _ = fi25
                        _ = v26
                    }
                    v27 := GenS2{F0: 42, F1: 928}
                    fmt.Println(v11, v1, v5)
                    _ = v27
                }
            } else {
                v5 = 1e10
            }
        }
        v5 -= 3.14159
    }
    fmt.Println(v5, v5, v11)
    var v28 uint16 = 489
    v1 = v1
    fmt.Println(v28)
    v29 := GenS3{F0: -2147483648, F1: 29, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
