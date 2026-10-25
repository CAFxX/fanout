package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
func genFn0(p0 string, p1 string, p2 bool) int8 {
    var v7 int64 = 9223372036854775807
    if (v7 <= (-(v7))) {
        if (482 < 2) {
            v8 := []int{521, 1}
            fmt.Println(p1, v7)
            _ = v8
        }
        v9 := []int{9223372036854775807, -1}
        _ = v9
    }
    var v10 uint16 = 0
    if p2 {
        v11 := []int{42, -9223372036854775808, 103, -637}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 0
}
func genFn1() bool {
    if ((-128 != -128) && (genFn0((("hello" + "foo bar") + "hello"), "", (true || true)) != 100)) {
        if true {
            var v12 int8 = 87
            if (true && (!false)) {
                var v13 uint = 65535
                var v14 uint8 = 1
                v12 = genFn0(("" + ("\n\t" + "")), (("" + (("a" + "a") + "\n\t")) + ("a" + "hello")), (0 == 2))
                _ = v13
                _ = v14
            }
            var v15 string = ((("foo bar" + "a") + "foo bar") + ("\n\t" + (("\n\t" + "\n\t") + "a")))
            _ = v12
            _ = v15
        }
        if (21 < 31) {
            var v16 int32 = -1
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v16)
            default:
                var v17 int = 0
                _ = v17
            }
            _ = v16
        }
    } else {
        if ((42 > 128) && true) {
            var v18 int64 = 1
            v18 -= v18
            _ = v18
        } else {
            var v19 string = ("foo bar" + "\n\t")
            var v20 bool = false
            var v21 uint16 = 2
            _ = v19
            _ = v20
            _ = v21
        }
        if ((("foo bar" + "\n\t") + ("a" + "")) > ("\n\t" + "")) {
            var v22 int8 = -1
            _ = v22
        }
        var v23 int = 0
        _ = v23
    }
    if (255 == 128) {
        var v24 uint = 65535
        v25 := GenS1{F0: 0, F1: (v24 + 65535)}
        _ = v24
        _ = v25
    } else {
        var v26 int = 42
        if false {
            switch "a" {
            case "hello":
                v27 := GenS1{F0: 0, F1: 4294967295}
                _ = v27
            case "\n\t":
                v26 = v26
            }
            fmt.Println(v26)
            fmt.Println(v26)
        }
        if (((("hello" + "") + (("a" + "\n\t") + "a")) + ("\n\t" + "\n\t")) <= "\n\t") {
            fmt.Println(v26)
        }
        _ = v26
    }
    v28 := []int{}
    var v29 int16 = 42
    _ = v28
    _ = v29
    return true
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]string)
    var v3 float32 = v1
    v4 := make(map[string]int)
    var v5 uint32 = 474
    var v6 string = "a"
    v1 = 3.14159
    fmt.Println(v3)
    if false {
        var v30 float32 = v3
        for fi31 := 0; fi31 < 11; fi31++ {
            v3 = v1
            _ = fi31
        }
        fmt.Println(v1, v1)
        _ = v30
    } else {
        var v32 uint64 = 2
        _ = v32
    }
    if (127 < 42) {
        var v33 string = v6
        _ = v33
    } else {
        var v34 uint = 1
        if false {
            fmt.Println(v34)
        } else {
            var v35 string = v6
            _ = v35
        }
        _ = v34
    }
    var v36 int16 = 139
    var v37 uint = 976
    switch (v6 + v6) {
    case "":
        v38 := []int{}
        _ = v38
    case "hello":
        switch v6 {
        case "hello":
            var v39 int32 = -106
            _ = v39
        case "a":
            for fi40 := 0; fi40 < 17; fi40++ {
                var v41 uint8 = 255
                var v42 bool = (255 >= v37)
                _ = fi40
                _ = v41
                _ = v42
            }
        }
    case "a":
        if false {
            var v43 float64 = 3.14159
            var v44 int8 = 100
            switch v6 {
            case "":
                switch v6 {
                case "a":
                    v45 := GenS2{F0: 1, F1: 0, F2: (((v6 + v6) + v6) + ((("\n\t" + "hello") + ("foo bar" + "\n\t")) + v6))}
                    _ = v45
                case "\n\t":
                    if true {
                        var v46 int64 = int64(v37)
                        v4["foo bar"] = 42
                        _ = v46
                    } else {
                        fmt.Println(v36, v5)
                    }
                case "":
                    v36 = 995
                }
            case "a":
                var v47 float32 = 1.5
                _ = v47
            }
            _ = v43
            _ = v44
        } else {
            v37 = (^(v37))
            v48 := v2[("foo bar" + v6)]
            var v49 int16 = v36
            _ = v48
            _ = v49
        }
    }
    v3 -= ((float32(1.5) * 0.0) - float32(float64(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
