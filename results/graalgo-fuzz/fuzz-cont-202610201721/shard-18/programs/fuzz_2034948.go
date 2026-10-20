package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int16) uint32 {
    if (1 >= 1) {
        var v6 string = ("\n\t" + "")
        var v7 string = v6
        switch -9223372036854775808 {
        case -988:
            var v8 int8 = -29
            _ = v8
        case 9223372036854775807:
            fmt.Println(v7, v6)
        }
        _ = v6
        _ = v7
    } else {
        var v9 uint32 = 750
        switch (("" + "foo bar") + ((("" + ("" + "hello")) + (("\n\t" + "\n\t") + "a")) + ("a" + "a"))) {
        case "\n\t":
            v9 -= v9
        case "hello":
            var v10 uint64 = 42
            _ = v10
        default:
            v9 = 4294967295
        }
        _ = v9
    }
    var v11 int32 = 2147483647
    var v12 uint8 = 191
    _ = p0
    _ = v11
    _ = v12
    return 178
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    v3 := []string{("hello" + "foo bar"), "foo bar", ("foo bar" + (((("hello" + ("foo bar" + ("hello" + "a"))) + "a") + "hello") + "foo bar")), "foo bar"}
    var v4 float32 = 1e10
    v5 := make(map[int]string)
    v3 = append(v3, "foo bar")
    fmt.Println(v4)
    v4 += 3.14159
    fmt.Println(v4)
    v13 := v2["hello"]
    switch -444 {
    case -603:
        for fi14 := 0; fi14 < 18; fi14++ {
            if false {
                v15 := GenS1{F0: 31}
                v13 = v13
                _ = v15
            } else {
                var v16 int32 = 0
                var v17 float32 = v4
                var v18 uint = 804
                _ = v16
                _ = v17
                _ = v18
            }
            var v19 bool = v13
            var v20 int8 = 1
            _ = fi14
            _ = v19
            _ = v20
        }
    case 936:
        if v13 {
            var v21 uint = 65535
            var v22 int32 = 42
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 14; fi23++ {
                fmt.Println(fi23)
                _ = fi23
            }
        }
    case -2147483648:
        v13 = v13
    }
    v24 := GenS1{F0: 2}
    if (!(-2.25 >= float32((-2.25 + ((-2.25 / 1e10) * 0.0))))) {
        var v25 int8 = -1
        var v26 uint = 4294967295
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        switch fi27 {
        case -978:
            v3 = append(v3, "")
        case 128:
            if v13 {
                v28 := v2[(("\n\t" + "foo bar") + (("a" + "\n\t") + "a"))]
                _ = v28
            } else {
                var v29 uint16 = 42
                fmt.Println(fi27, v4)
                v30 := GenS1{F0: 42}
                _ = v29
                _ = v30
            }
        case 127:
            switch (fi27 + fi27) {
            case -631:
                v4 = v4
            default:
                _ = len(v5)
            }
        }
        v4 -= v4
        v13 = v13
        _ = fi27
    }
    v13 = (("foo bar" + "\n\t") == ("foo bar" + "\n\t"))
    fmt.Println(v13, v13)
    for fi31 := 0; fi31 < 14; fi31++ {
        switch (fi31 >> ((1) & 63)) {
        case -354:
            for fi32 := 0; fi32 < 2; fi32++ {
                fmt.Println(v4)
                _ = fi32
            }
        default:
            var v33 bool = v13
            _ = v33
        }
        fmt.Println(v4, fi31)
        _ = fi31
    }
    for fi34 := 0; fi34 < 4; fi34++ {
        v13 = (4294967295 >= 0)
        v35 := v1[(fi34 << ((2) & 63))]
        v35 = v35
        _ = fi34
        _ = v35
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        v37 := GenS1{F0: 231}
        switch ("a" + "a") {
        case "foo bar":
            var v38 uint32 = 42
            _ = v38
        }
        var v39 int64 = 42
        _ = fi36
        _ = v37
        _ = v39
    }
    switch ("foo bar" + (("foo bar" + ("foo bar" + "")) + ("hello" + "hello"))) {
    case "\n\t":
        v40 := GenS1{F0: 2}
        _ = v40
    case "foo bar":
        var v41 uint16 = 65535
        _ = v41
    case "a":
        v13 = v13
    default:
        if v13 {
            v13 = true
        }
    }
    var v42 uint64 = 1
    var v43 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
