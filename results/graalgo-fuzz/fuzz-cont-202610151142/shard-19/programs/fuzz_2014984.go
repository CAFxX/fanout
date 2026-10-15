package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 uint; F2 string }
func genFn0(p0 string) int {
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int64 = 42
        for fi9 := 0; fi9 < 20; fi9++ {
            if (0 < 127) {
                v10 := GenS1{F0: 0}
                switch p0 {
                case "a":
                    v11 := []int{(fi9 & (^(fi7))), fi9}
                    _ = v11
                case "":
                    p0 = "hello"
                case "hello":
                    if (991 != 0) {
                        var v12 int8 = 0
                        _ = v12
                    } else {
                        v13 := []int{(-(fi7)), 32768}
                        _ = v13
                    }
                default:
                    v8 = (^(v8))
                }
                _ = v10
            } else {
                v14 := []int{32768, fi9, fi7, (fi7 | fi7)}
                v15 := GenS2{F0: 42}
                _ = v14
                _ = v15
            }
            v16 := []int{fi9, fi7, fi7}
            _ = fi9
            _ = v16
        }
        _ = fi7
        _ = v8
    }
    switch 1 {
    case 32767:
        if false {
            switch ((((p0 + p0) + "\n\t") + "foo bar") + ((p0 + "hello") + "\n\t")) {
            case "\n\t":
                if ((true && (p0 > p0)) || true) {
                    v17 := []int{372, 4294967295, 4294967295, -9223372036854775808}
                    v17 = append(v17, 128)
                    v17 = append(v17, 608)
                    _ = v17
                } else {
                    p0 = (p0 + "\n\t")
                }
            }
            if true {
                v18 := GenS1{F0: 0}
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 int16 = 1000
                    fmt.Println(p0)
                    _ = fi19
                    _ = v20
                }
                switch ("hello" + "foo bar") {
                case "foo bar":
                    v21 := GenS2{F0: 255}
                    _ = v21
                }
                _ = v18
            }
            fmt.Println(p0)
        } else {
            p0 = p0
            fmt.Println(p0)
        }
    }
    _ = p0
    return 1
}

func main() {
    v1 := []string{("foo bar" + "foo bar"), (((("foo bar" + ("hello" + "a")) + ("hello" + ("hello" + "a"))) + ("" + "foo bar")) + "a"), "\n\t", ("\n\t" + "hello")}
    var v2 int8 = -93
    v3 := []string{"a"}
    var v4 int8 = (v2 / ((v2) | 1))
    var v5 int8 = 100
    v6 := []uint32{4294967295, 65535}
    var v22 int8 = (v5 << ((74) & 63))
    var v23 int = -2147483648
    if (((((0.1 - 3.14159) / 1.5) <= 3.14159) || (!false)) || true) {
        if (("foo bar" != "foo bar") || (!(!true))) {
            fmt.Println(v23)
            for fi24 := 0; fi24 < 15; fi24++ {
                var v25 float32 = 1e10
                v2 = v4
                v6 = append(v6, 2)
                _ = fi24
                _ = v25
            }
        }
        if true {
            v4 -= 0
            v26 := GenS2{F0: 217}
            _ = v26
        }
    } else {
        v27 := GenS2{F0: 255}
        v23 = (v23 - (v23 + (v23 & 32767)))
        var v28 uint32 = 722
        _ = v27
        _ = v28
    }
    switch "\n\t" {
    case "hello":
        v22 -= v5
    case "":
        var v29 int64 = -301
        _ = v29
    case "foo bar":
        for fi30 := 0; fi30 < 3; fi30++ {
            var v31 uint = 2
            v32 := GenS1{F0: 0}
            switch "\n\t" {
            case "":
                for fi33 := 0; fi33 < 16; fi33++ {
                    var v34 int32 = 2147483647
                    var v35 string = "foo bar"
                    v4 = (v5 | (-(v2)))
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            default:
                switch "\n\t" {
                case "a":
                    fmt.Println(v22)
                }
            }
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    if true {
        v36 := GenS3{F0: 42, F1: v23, F2: 824}
        for fi37 := 0; fi37 < 16; fi37++ {
            switch "" {
            case "foo bar":
                v22 += (v5 % (((v2 & -128)) | 1))
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "":
                v4 = (v5 % ((1) | 1))
            }
            switch 127 {
            case 99:
                v6 = append(v6, 2147483647)
            case 279:
                _ = len(v3)
            }
            _ = fi37
        }
        switch (("\n\t" + "foo bar") + "") {
        case "a":
            if true {
                v38 := GenS4{F0: 913}
                _ = v38
            }
        case "":
            v39 := GenS4{F0: 2}
            _ = v39
        }
        _ = v36
    } else {
        if true {
            v40 := GenS5{F0: v22, F1: 18446744073709551615, F2: "hello"}
            _ = v40
        }
    }
    v3 = append(v3, (("" + (((("" + (("foo bar" + "foo bar") + "a")) + "\n\t") + "a") + ("" + ("hello" + (("" + "foo bar") + "\n\t"))))) + "a"))
    fmt.Println(v23, v23, v4)
    for fi41 := 0; fi41 < 7; fi41++ {
        if true {
            v22 = (^(v4))
            _ = len(v6)
            var v42 uint8 = 255
            _ = v42
        } else {
            var v43 string = "foo bar"
            var v44 uint16 = 255
            _ = v43
            _ = v44
        }
        fmt.Println(v5, v2)
        var v45 uint8 = 42
        _ = fi41
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
