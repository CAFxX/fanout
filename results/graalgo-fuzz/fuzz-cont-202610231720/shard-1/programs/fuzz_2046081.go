package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 255
    var v3 string = (((("\n\t" + (("foo bar" + "hello") + "")) + (("a" + (("\n\t" + ("hello" + "foo bar")) + "a")) + "\n\t")) + "hello") + "foo bar")
    switch v3 {
    case "foo bar":
        var v4 float64 = 1e10
        _ = v4
    case "\n\t":
        var v5 uint = 65535
        _ = v5
    default:
        var v6 string = v3
        _ = v6
    }
    v2 -= 2
    var v7 uint16 = (^(v2))
    if true {
        switch -985 {
        case 670:
            switch 0 {
            case 4294967295:
                var v8 bool = (1 >= 1)
                _ = v8
            case 42:
                v9 := GenS1{F0: 1}
                _ = v9
            case 9223372036854775807:
                v10 := GenS1{F0: 797}
                _ = v10
            }
        default:
            var v11 int32 = -2147483648
            _ = v11
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int64 = 314
            v14 := GenS2{F0: 99}
            var v15 int64 = (v13 << ((428) & 63))
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        v3 = ""
    } else {
        v16 := GenS3{F0: 18446744073709551615, F1: false, F2: 128}
        var v17 int16 = 1000
        _ = v16
        _ = v17
    }
    if (!false) {
        if true {
            delete(v1, 2147483648)
            fmt.Println(v2, v3, v2)
        } else {
            if false {
                v18 := GenS3{F0: 18446744073709551615, F1: (!(true || (true || false))), F2: 1}
                for fi19 := 0; fi19 < 9; fi19++ {
                    fmt.Println(v2, v7)
                    switch v3 {
                    case "foo bar":
                        for fi20 := 0; fi20 < 10; fi20++ {
                            v3 = v3
                            var v21 float32 = (3.14159 + 1e10)
                            v7 = (v2 & (v7 + v2))
                            _ = fi20
                            _ = v21
                        }
                    case "":
                        v7 = 255
                    case "\n\t":
                        var v22 int8 = -39
                        _ = v22
                    default:
                        v23 := GenS2{F0: 0}
                        _ = v23
                    }
                    v3 = "hello"
                    _ = fi19
                }
                var v24 float64 = 3.14159
                _ = v18
                _ = v24
            }
            v1[42] = (700 == 262)
        }
    }
    v2 -= v2
    var v25 float32 = 0.1
    for fi26 := 0; fi26 < 18; fi26++ {
        switch (fi26 >> ((18446744073709551615) & 63)) {
        case 4294967295:
            v25 = v25
        default:
            v27 := []int{fi26, fi26, -967}
            _ = v27
        }
        fmt.Println(v25)
        _ = fi26
    }
    fmt.Println(v3, v7)
    for fi28 := 0; fi28 < 1; fi28++ {
        fmt.Println(v7, fi28, v2)
        if (231 > 128) {
            v3 = v3
        }
        _ = fi28
    }
    switch (v3 + v3) {
    case "":
        var v29 uint8 = 2
        _ = v29
    case "foo bar":
        v7 -= (v2 >> ((4294967295) & 63))
    }
    v30 := []int{171}
    switch -9223372036854775808 {
    case 4294967295:
        delete(v1, 65535)
    case 32768:
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v7
}
