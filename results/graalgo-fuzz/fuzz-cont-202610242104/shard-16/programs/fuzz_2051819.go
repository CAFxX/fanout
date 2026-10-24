package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 string }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int8 }
func genFn0() int16 {
    var v5 string = ("foo bar" + "\n\t")
    for fi6 := 0; fi6 < 1; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            switch fi6 {
            case 127:
                var v8 int64 = 9223372036854775807
                _ = v8
            }
            _ = fi7
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 uint8 = 2
        v10 = (v10 & 127)
        v11 := GenS1{F0: (^(fi9)), F1: 2147483647, F2: v5}
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = 32767
    fmt.Println(v5)
    _ = v12
    _ = v5
    return -452
}

func main() {
    var v1 bool = ((4294967295 >= 127) && true)
    v2 := []string{"a"}
    v3 := make(map[string]uint32)
    var v4 float64 = 0.0
    for fi13 := 0; fi13 < 7; fi13++ {
        if v1 {
            var v14 uint = 42
            v15 := GenS2{F0: -9223372036854775808, F1: 2}
            _ = v14
            _ = v15
        } else {
            var v16 string = "\n\t"
            if len(v2) > 0 {
                v2[0] = ""
            }
            if (0 >= 32767) {
                var v17 int8 = 20
                _ = v17
            }
            _ = v16
        }
        _ = fi13
    }
    _ = len(v3)
    switch 128 {
    case 564:
        if (((((0.0 - 1.5) + (v4 * float64(float32(-2.25)))) + 0.1) + float64(1.5)) > float64(0.1)) {
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 string = "foo bar"
                _ = fi18
                _ = v19
            }
            var v20 uint64 = 902
            _ = v20
        }
    case -541:
        if len(v2) > 0 {
            v2[0] = (("hello" + "hello") + "foo bar")
        }
    default:
        v1 = (686 < 2147483647)
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        fmt.Println(fi21)
        switch "" {
        case "hello":
            if (2 < 0) {
                var v22 uint = 255
                var v23 uint64 = 792
                _ = v22
                _ = v23
            }
        case "":
            var v24 int32 = 0
            _ = v24
        case "\n\t":
            switch (("\n\t" + "hello") + ("a" + ("\n\t" + ("" + "")))) {
            case "":
                if v1 {
                    v4 = v4
                    switch "foo bar" {
                    case "foo bar":
                        v25 := v3[("hello" + "a")]
                        _ = v25
                    case "":
                        switch "hello" {
                        case "hello":
                            delete(v3, (("hello" + "foo bar") + "a"))
                        }
                    }
                    var v26 string = "hello"
                    _ = v26
                } else {
                    fmt.Println(v1, v4)
                    v27 := GenS1{F0: fi21, F1: 2147483648, F2: ("hello" + "")}
                    if (!v1) {
                        v1 = ("" >= "")
                    } else {
                        switch "hello" {
                        case "hello":
                            v4 -= v4
                        case "":
                            v4 = 1.5
                        }
                    }
                    _ = v27
                }
            }
        }
        _ = fi21
    }
    var v28 float64 = 1e10
    var v29 int16 = genFn0()
    v30 := GenS2{F0: 655, F1: 698}
    for fi31 := 0; fi31 < 9; fi31++ {
        for fi32 := 0; fi32 < 15; fi32++ {
            v3[("a" + ("a" + "a"))] = 1
            v4 -= v28
            _ = fi32
        }
        _ = fi31
    }
    var v33 float32 = 1e10
    v34 := GenS1{F0: 42, F1: 2147483647, F2: "hello"}
    var v35 int8 = -128
    fmt.Println(v28)
    if (("foo bar" >= "") || true) {
        for fi36 := 0; fi36 < 16; fi36++ {
            if ((v35 & v35) > (v35 << ((42) & 63))) {
                switch (fi36 / ((int(v29)) | 1)) {
                case 4294967295:
                    v3[("hello" + ("hello" + ("" + "hello")))] = 255
                case -1:
                    switch fi36 {
                    case -9223372036854775808:
                        var v37 float32 = v33
                        _ = v37
                    }
                case -9223372036854775808:
                    var v38 uint = 255
                    _ = v38
                default:
                    var v39 bool = (true && (-411 >= 42))
                    _ = v39
                }
            }
            var v40 int8 = v35
            _ = fi36
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 9; fi41++ {
            for fi42 := 0; fi42 < 17; fi42++ {
                var v43 uint = 4294967295
                fmt.Println(v4, v43, v4)
                v44 := GenS3{F0: -1}
                _ = fi42
                _ = v43
                _ = v44
            }
            _ = fi41
        }
        v45 := GenS4{F0: 18446744073709551615, F1: (v35 * v35)}
        _ = v45
    }
    v35 += v35
    v33 += 0.0
    var v46 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v46
}
