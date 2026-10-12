package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() string {
    for fi5 := 0; fi5 < 6; fi5++ {
        switch "a" {
        case "hello":
            var v6 int64 = 42
            _ = v6
        case "\n\t":
            v7 := GenS1{F0: 0}
            _ = v7
        default:
            var v8 uint8 = 210
            _ = v8
        }
        _ = fi5
    }
    v9 := []int{1, 42}
    var v10 float32 = (0.1 + 0.0)
    var v11 uint16 = 32767
    _ = v10
    _ = v11
    _ = v9
    return "hello"
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    v1 = 42
    var v12 uint16 = 0
    v1 = 9223372036854775807
    var v13 string = ("" + "foo bar")
    var v14 uint = 167
    v13 = "\n\t"
    v13 = ((v13 + v13) + ("" + v13))
    var v15 int32 = 42
    if (((3.14159 - 0.1) + -2.25) == (0.1 * 1.5)) {
        if (false || true) {
            fmt.Println(v12, v14, v12)
        }
    }
    v16 := []int{-235}
    switch v13 {
    case "":
        if false {
            if (uint32(v12) >= 2147483647) {
                v17 := GenS1{F0: (v1 / (((v1 | v1)) | 1))}
                _ = v17
            }
        } else {
            v16 = append(v16, 127)
            switch -874 {
            case 908:
                v13 = "a"
            case 32767:
                var v18 string = v13
                _ = v18
            }
            var v19 uint8 = 1
            _ = v19
        }
    default:
        var v20 float64 = 1e10
        _ = v20
    }
    switch v13 {
    case "a":
        v13 = v13
    case "\n\t":
        switch 9223372036854775807 {
        case -115:
            var v21 uint16 = v12
            _ = v21
        case 32767:
            var v22 uint = uint(v12)
            _ = v22
        case 1:
            var v23 uint32 = 65535
            _ = v23
        }
    case "":
        var v24 uint = v14
        _ = v24
    default:
        v25 := GenS1{F0: (v1 / (((v1 & (v1 << (((v14 << (((v14 >> ((66) & 63))) & 63))) & 63)))) | 1))}
        _ = v25
    }
    switch (v13 + v13) {
    case "\n\t":
        v1 = v1
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        switch fi26 {
        case 9223372036854775807:
            for fi27 := 0; fi27 < 5; fi27++ {
                v28 := v3[v13]
                if true {
                    var v29 uint64 = 2
                    _ = v29
                }
                var v30 string = "foo bar"
                _ = fi27
                _ = v28
                _ = v30
            }
        case 32767:
            if (!(!(!false))) {
                v15 += (v15 * v15)
                if (true && ((v12 | (v12 + (v12 << (((v14 << (((v14 + v14)) & 63))) & 63)))) != 255)) {
                    var v31 string = (v13 + ("a" + (v13 + (v13 + v13))))
                    v12 = v12
                    var v32 uint = (v14 / ((48) | 1))
                    _ = v31
                    _ = v32
                }
                v33 := GenS1{F0: 575}
                _ = v33
            } else {
                v34 := GenS1{F0: (v1 * 773)}
                var v35 uint64 = 18446744073709551615
                _ = v34
                _ = v35
            }
        default:
            for fi36 := 0; fi36 < 7; fi36++ {
                if true {
                    v37 := GenS1{F0: (v1 + (v1 & 42))}
                    fmt.Println(v1, fi26)
                    fmt.Println(fi36, fi26)
                    v38 := v3[v13]
                    _ = v37
                    _ = v38
                }
                _ = fi36
            }
        }
        switch "a" {
        case "":
            var v39 uint = 65535
            _ = v39
        case "\n\t":
            switch v13 {
            case "\n\t":
                v12 = v12
            case "":
                fmt.Println(fi26, v14)
            case "foo bar":
                _ = len(v2)
            default:
                var v40 int = (fi26 ^ int(v14))
                _ = v40
            }
        }
        if false {
            v14 -= 61
            fmt.Println(fi26, v13)
            v3[v13] = v13
        } else {
            if true {
                v41 := GenS1{F0: v1}
                _ = v41
            } else {
                v13 = ((v13 + ("\n\t" + v13)) + ((("hello" + "foo bar") + (v13 + v13)) + ((v13 + "\n\t") + (v13 + ("foo bar" + v13)))))
            }
            v12 -= uint16(v14)
            switch (fi26 | (fi26 % (((fi26 << ((v14) & 63))) | 1))) {
            case 1:
                var v42 uint64 = 4294967295
                _ = v42
            case 128:
                _ = len(v16)
            case -252:
                for fi43 := 0; fi43 < 20; fi43++ {
                    var v44 string = (("\n\t" + v13) + v13)
                    switch v13 {
                    case "":
                        var v45 int32 = (v15 >> ((634) & 63))
                        v14 -= (v14 / ((v14) | 1))
                        _ = v45
                    }
                    _ = fi43
                    _ = v44
                }
            }
        }
        _ = fi26
    }
    v46 := GenS1{F0: v1}
    v47 := GenS1{F0: (v1 >> ((v14) & 63))}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
}
