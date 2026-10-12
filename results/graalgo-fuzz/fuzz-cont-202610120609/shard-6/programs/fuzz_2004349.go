package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
func genFn0(p0 int32, p1 string) uint32 {
    var v7 string = ""
    var v8 string = (v7 + ((((v7 + ((("hello" + "") + p1) + v7)) + (p1 + v7)) + p1) + ""))
    for fi9 := 0; fi9 < 3; fi9++ {
        p0 = (p0 + p0)
        for fi10 := 0; fi10 < 12; fi10++ {
            p1 = v7
            fmt.Println(p1)
            var v11 bool = true
            _ = fi10
            _ = v11
        }
        switch v8 {
        case "hello":
            v8 = ""
        case "\n\t":
            v12 := []int{-461, (^(fi9))}
            _ = v12
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]int)
    var v3 int16 = 42
    v4 := make(map[int]bool)
    var v5 uint64 = 395
    var v6 int64 = -557
    if false {
        v13 := GenS1{F0: v5}
        var v14 float64 = v1
        switch "hello" {
        case "foo bar":
            if true {
                if (("foo bar" + "\n\t") > ("\n\t" + ("" + "\n\t"))) {
                    delete(v2, (("" + "hello") + "foo bar"))
                    v15 := []int{1, 2147483647, 127}
                    for fi16 := 0; fi16 < 1; fi16++ {
                        var v17 uint = 65535
                        v5 += (v5 / (((v5 + v5)) | 1))
                        var v18 int = fi16
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    _ = v15
                } else {
                    var v19 string = (("a" + "hello") + "hello")
                    var v20 uint = 2
                    _ = v19
                    _ = v20
                }
                v6 -= int64(v3)
                var v21 uint8 = 53
                _ = v21
            } else {
                if false {
                    v22 := []int{32768, 127, 127}
                    _ = v22
                }
            }
        case "a":
            var v23 bool = false
            _ = v23
        case "hello":
            for fi24 := 0; fi24 < 15; fi24++ {
                delete(v4, 32768)
                v1 = (v14 + 1.5)
                v3 += 167
                _ = fi24
            }
        }
        _ = v13
        _ = v14
    }
    var v25 int16 = v3
    v3 += v3
    fmt.Println(v5)
    if true {
        if (!((v1 / v1) != (float64(0.0) + ((v1 * v1) - float64(0.0))))) {
            v26 := []int{0}
            v27 := GenS2{F0: 65535}
            _ = v26
            _ = v27
        }
        switch (("\n\t" + "a") + "") {
        case "\n\t":
            var v28 float64 = 3.14159
            _ = v28
        case "a":
            v29 := []int{}
            _ = v29
        }
    }
    var v30 float64 = -2.25
    var v31 uint64 = v5
    for fi32 := 0; fi32 < 18; fi32++ {
        for fi33 := 0; fi33 < 11; fi33++ {
            v34 := []int{(fi32 >> ((255) & 63)), (fi33 ^ fi32), (-(fi32)), (fi32 << ((4294967295) & 63))}
            if (!(507 == genFn0(1, "a"))) {
                switch fi33 {
                case 127:
                    fmt.Println(v3)
                case 138:
                    var v35 float64 = -2.25
                    _ = v35
                }
            }
            v31 = (v31 << ((0) & 63))
            _ = fi33
            _ = v34
        }
        _ = len(v2)
        switch "hello" {
        case "hello":
            if (!(!true)) {
                v6 = 9223372036854775807
                v6 = v6
            }
        case "foo bar":
            v36 := GenS1{F0: v31}
            _ = v36
        default:
            var v37 uint = 4294967295
            _ = v37
        }
        _ = fi32
    }
    v38 := v2[""]
    v39 := GenS1{F0: v31}
    v40 := GenS2{F0: 255}
    for fi41 := 0; fi41 < 20; fi41++ {
        var v42 uint16 = 255
        v43 := []int{v38, -307, fi41}
        _ = fi41
        _ = v42
        _ = v43
    }
    v25 -= (v3 << ((255) & 63))
    for fi44 := 0; fi44 < 7; fi44++ {
        v30 += (-2.25 * v1)
        for fi45 := 0; fi45 < 2; fi45++ {
            var v46 string = ("" + "hello")
            _ = fi45
            _ = v46
        }
        for fi47 := 0; fi47 < 13; fi47++ {
            if (v6 <= v6) {
                switch "\n\t" {
                case "":
                    v30 = (float64(-2.25) + v1)
                case "a":
                    v2["\n\t"] = (fi47 % ((128) | 1))
                case "hello":
                    var v48 int = (v38 ^ (v38 / ((fi47) | 1)))
                    _ = v48
                default:
                    v49 := []int{}
                    _ = v49
                }
                v50 := []int{911, -1, fi44}
                _ = v50
            }
            _ = fi47
        }
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
