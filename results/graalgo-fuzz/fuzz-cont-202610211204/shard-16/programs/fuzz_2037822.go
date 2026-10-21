package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0() string {
    var v5 int64 = 837
    fmt.Println(v5)
    v6 := GenS1{F0: "\n\t"}
    var v7 int32 = 0
    var v8 float32 = 0.1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (("hello" + ((("a" + "a") + ("\n\t" + "foo bar")) + ("foo bar" + "\n\t"))) + "foo bar")
}
func genFn1() uint32 {
    var v9 int8 = 42
    v10 := []int{-547}
    _ = v10
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint16 = 0
    var v2 uint32 = 384
    var v3 int8 = -1
    var v4 uint32 = (v2 << ((255) & 63))
    var v11 int64 = 193
    v12 := GenS1{F0: "a"}
    if (!true) {
        v13 := []int{2147483647, 992, 945, 42}
        for fi14 := 0; fi14 < 17; fi14++ {
            if false {
                var v15 uint = 1
                _ = v15
            } else {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                switch fi14 {
                case -87:
                    var v16 int16 = -32768
                    _ = v16
                case -92:
                    v3 = (v3 & v3)
                default:
                    var v17 uint32 = v2
                    _ = v17
                }
                if true {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        v19 := GenS1{F0: "\n\t"}
                        fmt.Println(v4, v1, v3)
                        v20 := GenS1{F0: "foo bar"}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    var v21 int8 = int8(v1)
                    _ = v21
                }
            }
            if (v1 <= (v1 - v1)) {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            } else {
                var v22 uint8 = 2
                if (false || (((v11 <= (v11 % (((v11 - v11)) | 1))) && true) || false)) {
                    v23 := GenS1{F0: ""}
                    switch "" {
                    case "\n\t":
                        var v24 uint8 = uint8(v1)
                        _ = v24
                    case "":
                        v3 = (v3 << ((1) & 63))
                    }
                    v25 := GenS1{F0: ((("a" + "hello") + "hello") + "\n\t")}
                    _ = v23
                    _ = v25
                }
                _ = v22
            }
            _ = fi14
        }
        _ = v13
    }
    v26 := []int{-687}
    var v27 string = "foo bar"
    v28 := GenS1{F0: v27}
    switch ("\n\t" + (v27 + "hello")) {
    case "a":
        var v29 float32 = 3.14159
        _ = v29
    }
    if len(v26) > 0 {
        v26[0] = -2147483648
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            v2 = genFn1()
            var v32 float32 = -2.25
            switch (^(fi31)) {
            case 42:
                var v33 uint = 18446744073709551615
                _ = v33
            case 127:
                var v34 float32 = 1e10
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        _ = fi30
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        v4 = 2
        _ = fi35
    }
    if (-959 == 1) {
        v11 -= v11
        v36 := GenS1{F0: ((v27 + v27) + (((v27 + v27) + v27) + (v27 + v27)))}
        _ = v36
    }
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    var v37 int8 = (v3 + (v3 / (((v3 + (v3 ^ -1))) | 1)))
    v4 = v4
    if ((52 == 1) || (!true)) {
        var v38 int = 42
        for fi39 := 0; fi39 < 12; fi39++ {
            switch v27 {
            case "hello":
                var v40 float32 = 0.0
                _ = v40
            default:
                switch ((v27 + "foo bar") + "hello") {
                case "a":
                    v2 = (v4 & 2147483648)
                case "hello":
                    var v41 int8 = v3
                    _ = v41
                case "foo bar":
                    if true {
                        v42 := GenS2{F0: -1, F1: 5}
                        var v43 uint16 = (v1 & (v1 - (v1 ^ 0)))
                        var v44 float64 = 1.5
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                }
            }
            _ = fi39
        }
        _ = v38
    } else {
        _ = len(v26)
        v4 = v2
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v4
}
