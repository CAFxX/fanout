package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 bool = (776 > 0)
    var v2 int8 = 100
    var v3 int = 9223372036854775807
    v4 := make(map[uint32]uint32)
    for fi5 := 0; fi5 < 10; fi5++ {
        v2 += v2
        v2 -= (^(v2))
        _ = fi5
    }
    if v1 {
        switch (v3 % (((v3 | v3)) | 1)) {
        case 127:
            fmt.Println(v2)
        case 128:
            v2 -= 58
        case 32767:
            v2 += (v2 / ((127) | 1))
        }
        v6 := GenS1{F0: "hello"}
        if ("foo bar" > "hello") {
            v7 := GenS1{F0: "hello"}
            v8 := GenS1{F0: "hello"}
            _ = v7
            _ = v8
        } else {
            if (!((0.0 * 0.0) < 0.1)) {
                switch (v3 - v3) {
                case 32767:
                    var v9 uint8 = 42
                    _ = v9
                default:
                    v10 := []int{}
                    _ = v10
                }
                v3 = v3
            } else {
                if ((0 <= 186) || (42 == 1)) {
                    var v11 bool = (963 == 65535)
                    fmt.Println(v1, v2, v3)
                    var v12 bool = false
                    _ = v11
                    _ = v12
                }
            }
            var v13 uint16 = 32767
            _ = v13
        }
        _ = v6
    }
    if (144 > -1) {
        fmt.Println(v2)
        switch 32768 {
        case 928:
            switch (v3 % ((v3) | 1)) {
            case -2147483648:
                var v14 uint32 = 2147483647
                _ = v14
            }
        }
        v15 := []int{v3, (v3 | (^(v3))), (-(v3)), (v3 ^ 32767)}
        _ = v15
    } else {
        v16 := []int{}
        var v17 uint32 = 0
        if true {
            if (908 < 2147483647) {
                var v18 int16 = 42
                _ = v18
            }
        }
        _ = v16
        _ = v17
    }
    if (!(v1 && v1)) {
        var v19 uint32 = 355
        _ = v19
    }
    var v20 float64 = 1.5
    v21 := GenS1{F0: (("\n\t" + "\n\t") + "a")}
    v22 := GenS1{F0: "hello"}
    v23 := v4[674]
    for fi24 := 0; fi24 < 12; fi24++ {
        switch ("a" + "\n\t") {
        case "":
            var v25 string = "\n\t"
            _ = v25
        case "a":
            v20 -= (float64((float32((((v20 / ((1e10 / (1e10 - v20)) + ((v20 * 1.5) - 0.1))) - v20) / ((v20 + v20) * 3.14159))) * ((-2.25 - 3.14159) - (0.1 + ((1e10 / 0.1) / (0.0 + ((0.1 + (0.0 * 0.0)) * (1e10 - (float32((v20 / v20)) - 0.1))))))))) / 3.14159)
        case "hello":
            var v26 bool = v1
            _ = v26
        default:
            if (0.1 < (3.14159 * ((3.14159 / 0.1) - v20))) {
                var v27 int32 = 786
                _ = v27
            }
        }
        var v28 int8 = 100
        _ = fi24
        _ = v28
    }
    var v29 int64 = 1
    v1 = false
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 float32 = (1e10 * -2.25)
        _ = fi30
        _ = v31
    }
    if v1 {
        fmt.Println(v29, v3)
        v32 := []int{v3, v3, 128, 32767}
        fmt.Println(v1, v2, v29)
        _ = v32
    }
    v23 = v23
    for fi33 := 0; fi33 < 6; fi33++ {
        for fi34 := 0; fi34 < 14; fi34++ {
            delete(v4, v23)
            v2 -= (-(v2))
            var v35 int = (v3 + 2147483647)
            delete(v4, (v23 | v23))
            switch (-(v3)) {
            case -9223372036854775808:
                delete(v4, 4294967295)
            case 0:
                var v36 int32 = -471
                _ = v36
            case 127:
                v3 -= (fi33 / ((-314) | 1))
            }
            v37 := GenS1{F0: "\n\t"}
            _ = fi34
            _ = v35
            _ = v37
        }
        _ = fi33
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 uint16 = 514
        if ("\n\t" < (("a" + ("\n\t" + (("foo bar" + (("\n\t" + (("hello" + (("a" + "hello") + "foo bar")) + "")) + "\n\t")) + ("" + "")))) + ("" + "hello"))) {
            if (2 != 18446744073709551615) {
                v40 := GenS1{F0: "\n\t"}
                fmt.Println(fi38)
                _ = v40
            }
            v29 = v29
            switch ("foo bar" + "a") {
            case "":
                v2 = (-(v2))
            }
        }
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
}
