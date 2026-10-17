package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int32, p2 int64) uint64 {
    var v6 uint32 = (^(p0))
    if (-9223372036854775808 > -1) {
        var v7 uint8 = 127
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                switch "hello" {
                case "a":
                    fmt.Println(p2, p0, fi8)
                }
                p1 = p1
                var v10 int32 = p1
                p2 = 0
                var v11 uint16 = 362
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        v12 := GenS1{F0: p1, F1: (("foo bar" + ("\n\t" + "foo bar")) + (("\n\t" + ("a" + ("foo bar" + ("foo bar" + "hello")))) + ("\n\t" + ""))), F2: false}
        _ = v12
        _ = v7
    } else {
        switch ("foo bar" + (("\n\t" + "") + ("a" + "\n\t"))) {
        case "a":
            v13 := GenS2{F0: ""}
            _ = v13
        }
        v14 := GenS2{F0: ((("\n\t" + ("\n\t" + "foo bar")) + ("hello" + "a")) + "a")}
        _ = v14
    }
    v15 := GenS1{F0: (p1 / (((p1 & (^(p1)))) | 1)), F1: "", F2: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{"", "foo bar", "foo bar", ("foo bar" + "foo bar")}
    var v3 uint = 0
    v4 := make(map[int]uint32)
    var v5 string = "hello"
    v16 := GenS1{F0: 42, F1: v5, F2: false}
    switch 0 {
    case 405:
        var v17 float32 = 0.0
        _ = v17
    case -648:
        switch 32767 {
        case 32768:
            v18 := GenS1{F0: 42, F1: ((("foo bar" + v5) + v5) + "hello"), F2: (false || (3.14159 <= 0.0))}
            _ = v18
        default:
            if false {
                v3 = (v3 ^ 4294967295)
            }
        }
    case 127:
        v5 = v5
    default:
        fmt.Println(v3, v3)
    }
    var v19 int = 0
    switch 127 {
    case 2147483647:
        v3 = v3
    }
    switch 4294967295 {
    case -2147483648:
        if true {
            var v20 uint16 = 615
            for fi21 := 0; fi21 < 16; fi21++ {
                v5 = v5
                if true {
                    switch (^(fi21)) {
                    case 9223372036854775807:
                        fmt.Println(v20, v20)
                    case 32767:
                        var v22 int64 = 177
                        _ = v22
                    }
                    v23 := v1[-817]
                    _ = v23
                } else {
                    var v24 int32 = 42
                    v19 -= (^(fi21))
                    _ = v24
                }
                _ = fi21
            }
            _ = v20
        } else {
            if true {
                fmt.Println(v19)
                for fi25 := 0; fi25 < 19; fi25++ {
                    v2 = append(v2, v5)
                    fmt.Println(v5)
                    _ = fi25
                }
            }
            switch 42 {
            case 32767:
                var v26 int = (v19 | (v19 & 128))
                _ = v26
            case -2147483648:
                v27 := GenS3{F0: "\n\t", F1: false, F2: 32767}
                _ = v27
            default:
                v19 -= (v19 / (((v19 * v19)) | 1))
            }
            if len(v2) > 0 {
                v2[0] = ("hello" + "")
            }
        }
    case -1:
        fmt.Println(v3, v3, v19)
    }
    var v28 bool = false
    var v29 float64 = 1.5
    var v30 string = v5
    switch (v19 % ((-965) | 1)) {
    case -9223372036854775808:
        var v31 uint = 42
        _ = v31
    case 4294967295:
        var v32 float32 = 1.5
        _ = v32
    default:
        if ((32767 <= 42) && (-1 != -764)) {
            v5 = "foo bar"
            _ = len(v2)
            var v33 float64 = v29
            _ = v33
        } else {
            v3 = 65535
            var v34 string = v5
            _ = v34
        }
    }
    v35 := GenS1{F0: 2147483647, F1: v30, F2: ((!true) && (!v28))}
    if (42 == 0) {
        delete(v1, -580)
    }
    if v28 {
        if len(v2) > 0 {
            v2[0] = v5
        }
        for fi36 := 0; fi36 < 3; fi36++ {
            var v37 int16 = 0
            switch ("\n\t" + v5) {
            case "foo bar":
                var v38 int8 = 1
                _ = v38
            case "a":
                fmt.Println(fi36, v3)
            default:
                v39 := GenS4{F0: 429}
                _ = v39
            }
            _ = fi36
            _ = v37
        }
    }
    v40 := GenS2{F0: "a"}
    v29 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v40
    _ = v5
}
