package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
func genFn0(p0 int64, p1 int, p2 int32) int32 {
    var v4 int8 = 42
    for fi5 := 0; fi5 < 4; fi5++ {
        if false {
            v6 := []int{}
            _ = v6
        } else {
            var v7 string = "foo bar"
            _ = v7
        }
        for fi8 := 0; fi8 < 11; fi8++ {
            p1 -= (fi8 / (((fi5 * (fi5 | (fi8 % ((int(p2)) | 1))))) | 1))
            _ = fi8
        }
        _ = fi5
    }
    var v9 float64 = (float64(0.1) + -2.25)
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return p2
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[string]uint32)
    var v3 float32 = 1e10
    v3 = v3
    v3 -= ((v3 - 3.14159) * (v3 - v3))
    v10 := GenS1{F0: 42}
    v11 := []int{0}
    if (false && true) {
        v3 += 1.5
        var v12 int8 = 0
        _ = v12
    } else {
        v3 = (1.5 - v3)
        delete(v2, (("foo bar" + ((("foo bar" + ("a" + "\n\t")) + ("a" + (("a" + "") + "a"))) + "a")) + ("foo bar" + "")))
    }
    if (true || (1 <= 1)) {
        switch 4294967295 {
        case -9223372036854775808:
            for fi13 := 0; fi13 < 13; fi13++ {
                for fi14 := 0; fi14 < 4; fi14++ {
                    v15 := GenS1{F0: 0}
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
        case 128:
            switch ("a" + ("a" + (("a" + ("a" + "\n\t")) + "a"))) {
            case "foo bar":
                v16 := v2[""]
                _ = v16
            default:
                var v17 string = (("" + "a") + ("foo bar" + "a"))
                _ = v17
            }
        default:
            for fi18 := 0; fi18 < 3; fi18++ {
                v1 = v1
                switch (^(fi18)) {
                case -256:
                    var v19 int = (fi18 << ((1) & 63))
                    _ = v19
                case -9223372036854775808:
                    fmt.Println(v3, v1)
                default:
                    var v20 string = ("\n\t" + "\n\t")
                    _ = v20
                }
                _ = fi18
            }
        }
        v3 = v3
    } else {
        if (true || true) {
            var v21 uint = 1
            v3 += ((v3 / 3.14159) + (v3 - v3))
            var v22 uint16 = 2
            _ = v21
            _ = v22
        } else {
            var v23 uint32 = 2147483647
            _ = v23
        }
        if false {
            v3 = -2.25
            var v24 int8 = 127
            v25 := GenS2{F0: 4294967295}
            _ = v24
            _ = v25
        }
    }
    if (32767 >= 32768) {
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 uint64 = 2
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 int32 = 388
                _ = fi28
                _ = v29
            }
            _ = fi26
            _ = v27
        }
    } else {
        v30 := GenS3{F0: 914, F1: 1, F2: 65535}
        v3 = float32(3.14159)
        _ = v30
    }
    var v31 int64 = -9223372036854775808
    var v32 int64 = (-(v31))
    var v33 float32 = 1e10
    if (true || false) {
        v34 := GenS3{F0: 1, F1: 32767, F2: 650}
        _ = v34
    } else {
        var v35 int = -533
        v36 := GenS2{F0: 1}
        if (!(-110 <= -1)) {
            delete(v2, (("\n\t" + ("hello" + ("hello" + "\n\t"))) + "foo bar"))
        }
        _ = v35
        _ = v36
    }
    v3 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
